package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	layers "github.com/google/gopacket/layers"
)

var records map[string]string

func main() {
	fmt.Println("I Am Runihng")

	addr := net.UDPAddr{
		Port: 53,
		IP:   net.ParseIP("localhost"),
	}
	u, _ := net.ListenUDP("udp", &addr)

	// Wait to get request on that port
	for {
		tmp := make([]byte, 1024)
		_, addr, _ := u.ReadFrom(tmp)
		clientAddr := addr
		packet := gopacket.NewPacket(tmp, layers.LayerTypeDNS, gopacket.Default)
		dnsPacket := packet.Layer(layers.LayerTypeDNS)
		tcp, _ := dnsPacket.(*layers.DNS)
		serveDNS(u, clientAddr, tcp)
	}
}

func serveDNS(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	replyMess := request
	var dnsAnswer layers.DNSResourceRecord
	dnsAnswer.Type = layers.DNSTypeA
	var ip string
	var err error
	var ok bool
	ip, ok = records[string(request.Questions[0].Name)]
	if !ok {
		fmt.Println("Can't Find LOCALY: " + string(request.Questions[0].Name))
		ips, err := net.LookupIP(string(request.Questions[0].Name))
		if err != nil || len(ips) == 0 {
			dnsAnswer.Type = layers.DNSTypeA
			dnsAnswer.Name = []byte(request.Questions[0].Name)
			fmt.Println("Can't Find DNS GLOBALY!: " + string(request.Questions[0].Name))
			replyMess.ResponseCode = layers.DNSResponseCodeNXDomain
			dnsAnswer.Class = layers.DNSClassIN
			replyMess.QR = true
			replyMess.ANCount = 1
			replyMess.OpCode = layers.DNSOpCodeNotify
			replyMess.AA = true
			replyMess.Answers = append(replyMess.Answers, dnsAnswer)

		} else {
			fmt.Println("OK! (From Another DNS Server): " + string(request.Questions[0].Name))
			replyMess.QR = true
			replyMess.ANCount = 1
			replyMess.OpCode = layers.DNSOpCodeNotify
			replyMess.AA = true
			replyMess.ResponseCode = layers.DNSResponseCodeNoErr
			dnsAnswer.Name = []byte(request.Questions[0].Name)
			dnsAnswer.Type = layers.DNSTypeA
			dnsAnswer.Class = layers.DNSClassIN
			for _, i := range ips {
				dnsAnswer.IP = i
				replyMess.Answers = append(replyMess.Answers, dnsAnswer)
			}
		}

	} else {
		a, _, _ := net.ParseCIDR(ip + "/24")
		dnsAnswer.Type = layers.DNSTypeA
		dnsAnswer.IP = a
		dnsAnswer.Name = []byte(request.Questions[0].Name)
		fmt.Println("OK!: " + string(request.Questions[0].Name))
		replyMess.ResponseCode = layers.DNSResponseCodeNoErr
		dnsAnswer.Class = layers.DNSClassIN
		replyMess.QR = true
		replyMess.ANCount = 1
		replyMess.OpCode = layers.DNSOpCodeNotify
		replyMess.AA = true
		replyMess.Answers = append(replyMess.Answers, dnsAnswer)

	}

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}
	err = replyMess.SerializeTo(buf, opts)
	if err != nil {
		panic(err)
	}
	u.WriteTo(buf.Bytes(), clientAddr)
}
