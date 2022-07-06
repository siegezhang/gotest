package test

import (
	"flag"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func TestTcp(t *testing.T) {

	var srcip string
	var myport int
	flag.StringVar(&srcip, "ip", "", "source ip address")
	flag.IntVar(&myport, "port", -1, "my port to watch")
	flag.Parse()
	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	var b = make([]byte, 4096)
	for {
		n, addr, err := conn.ReadFrom(b)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if srcip != "" && addr.String() != srcip {
			continue
		}
		fmt.Println(n, " ", addr.String())
		packet := gopacket.NewPacket(b[:n], layers.LayerTypeTCP, gopacket.Default)
		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			if myport == -1 || (myport != -1 && myport == int(tcp.DstPort)) {
				fmt.Printf("SrcPort:%s\n", tcp.SrcPort.String())
				fmt.Printf("DstPort:%s\n", tcp.DstPort.String())
				fmt.Printf("Seq:%d\n", tcp.Seq)
				fmt.Printf("Ack:%d\n", tcp.Ack)
				fmt.Printf("URG:%t\n", tcp.URG)
				fmt.Printf("ACK:%t\n", tcp.ACK)
				fmt.Printf("PSH:%t\n", tcp.PSH)
				fmt.Printf("RST:%t\n", tcp.RST)
				fmt.Printf("SYN:%t\n", tcp.SYN)
				fmt.Printf("FIN:%t\n", tcp.FIN)
				fmt.Printf("Window:%d\n", tcp.Window)
				fmt.Printf("Checksum:%d\n", tcp.Checksum)
				fmt.Printf("Urgent:%d\n", tcp.Urgent)
				fmt.Printf("Options:%v\n", tcp.Options)
				fmt.Printf("Padding:%d\n", tcp.Padding)
				fmt.Printf("Contents:%d\n", tcp.Contents)
				fmt.Printf("LayerContents:%d\n", tcp.LayerContents())
			}
		}
	}
}
