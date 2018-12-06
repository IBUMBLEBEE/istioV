package main

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"

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

	router.Run(":9100")
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
		if len(device.Addresses) == 0 {
			continue
		}
		if device.Name == "docker0" || device.Name == "lo" {
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
			go func() {
				fmt.Println(packet)
			}()
		}
	}
}

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
