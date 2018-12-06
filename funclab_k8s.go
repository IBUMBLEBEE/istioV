package main

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/google/gopacket/pcap"
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
	go FindDevices()

	router.Run(":9180")
}

func FindDevices() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\n Name", device.Name)
		fmt.Println("Descripttion: ", device.Description)
		fmt.Println("Devices address: ", device.Addresses)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}

func GetGID() uint64 {
	// fmt.Println("The number of GetGID: ", runtime.NumGoroutine(), GetGID())
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
