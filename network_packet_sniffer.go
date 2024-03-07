/*
Don't forget to install below go library:
go get github.com/google/gopacket

Edit network interface and filter accordingly

And also install below:
---Ubuntu/Debian---
sudo apt-get install libpcap-dev

---Centos/RHEL---
sudo yum install libpcap-devel

---MacOS (via Homebrew)---
brew install libpcap
*/
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	// Define network interface and filter
	iface := "eth0"
	filter := "tcp"

	// Open network interface for packet capture
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set packet filter
	if err := handle.SetBPFFilter(filter); err != nil {
		log.Fatal(err)
	}

	// Capture packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Parse packet
		src := packet.NetworkLayer().NetworkFlow().Src().String()
		dst := packet.NetworkLayer().NetworkFlow().Dst().String()
		length := packet.Metadata().CaptureInfo.CaptureLength

		// Display packet information
		fmt.Printf("Src: %s, Dst: %s, Length: %d\n", src, dst, length)

		// Sleep for a short duration to avoid overwhelming the CPU
		time.Sleep(100 * time.Millisecond)
	}
}
