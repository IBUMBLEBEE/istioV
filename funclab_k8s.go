package main

import (
	"regexp"
	"bytes"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/google/gopacket/layers"

	"github.com/google/gopacket"

	"github.com/gin-gonic/gin"

	"github.com/google/gopacket/pcap"
)

var (
	snapshotLen int32 = 1024
	promiscuous       = false
	timeout           = 30 * time.Second
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
		fmt.Println("The number of ping: ", runtime.NumGoroutine(), GetGID())
	})

	router.GET("/num/goroutine", func(c *gin.Context) {
		go func() {
			fmt.Println("The number of goroutine: ", runtime.NumGoroutine(), GetGID())
		}()
	})

	// func group
	go OpenDeviceLiveCapture()

	router.Run(":9188")
}

// FindDevices find devices
func FindDevices() []string {
	deviceslice := []string{}
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Println(err)
	}

	// devicesliceaddr := &deviceslice
	// fmt.Println("Devices found:")
	for _, device := range devices {
		if len(device.Addresses) == 0 || device.Name == "docker0" || device.Name == "lo" {
			continue
		}
		// filter via network
		if device.Name == "flannel0" || device.Name == "dummy0" || device.Name == "kube-ipvs0" {
			continue
		}
		// filter cni
		if device.Name == "cni0" {
			continue
		}
		// filter veth device
		if (regexp.Match("veth.*", []byte(device.Name)) {
			continue
		}
		deviceslice = append(deviceslice, device.Name)
	}
	fmt.Println(deviceslice)
	return deviceslice
}

// OpenDeviceLiveCapture Capture packet
func OpenDeviceLiveCapture() {
	fmt.Println("OpenDeviceLiveCapture Runing...")
	devslice := FindDevices()
	for _, device := range devslice {
		handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
		if err != nil {
			log.Fatal(err)
		}

		// Use the handle as a packet source to process all packets
		packageSource := gopacket.NewPacketSource(handle, handle.LinkType())
		fmt.Println(packageSource)
		for packet := range packageSource.Packets() {
			go printPacketInfo(packet)
			// time.Sleep(5 * time.Second)
		}
	}
}

func printPacketInfo(packet gopacket.Packet) {
	fmt.Println("TcpTimestamp: ", time.Now().UnixNano())
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		fmt.Println("Ethernet layer detected")
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
		fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		fmt.Println()
	}

	// Let's see if the packet is IP (even though the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPV4 layer detected")
		ip, _ := ipLayer.(*layers.IPv4)

		// IP layer variables:
		// version (Either 4 or 6)
		// IHL (IP Header Length in 32-bit words)
		// TOS, Length, Id, Flags, FragOffset, TTL, Procotol (TCP?)
		// Checksum, srcIP, DstIP
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("protocol: ", ip.Protocol)
		fmt.Println()
	}

	// Let's see if the packet is TCP
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		fmt.Println("TCP layer detected")
		tcp, _ := tcpLayer.(*layers.TCP)

		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		fmt.Printf("From port %s to %d\n", tcp.SrcPort, tcp.DstPort)
		fmt.Println("Sequence number: ", tcp.Seq)
		fmt.Println()
	}

	// Iterate over all layer, printing out each layer type
	fmt.Println("All packet layers:")
	for _, layer := range packet.Layers() {
		fmt.Println("- ", layer.LayerType())
	}

	// When iterating through packet.Layers() above,
	// if it lists Payload layer then that is the same as
	// the applicationLayer. applicationLayer contains the payload
	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("Application layer/Payload found.")
		// fmt.Printf("%s\n", applicationLayer.Payload())

		// Search for a string inside the payload
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
			fmt.Println("HTTP found!")
		}
	}

	// Check for errors
	if err := packet.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet: ")
	}
	fmt.Println("========================================================================")
}

func Match(pattern string, b []byte) (matched bool, err error)


// GetGID get goroutine ID
func GetGID() uint64 {
	// fmt.Println("The number of GetGID: ", runtime.NumGoroutine(), GetGID())
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
