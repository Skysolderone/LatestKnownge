package main

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

func handler(writer dns.ResponseWriter, req *dns.Msg) {
	var resp dns.Msg
	resp.SetReply(req)
	for _, question := range req.Question {
		recordA := dns.A{
			Hdr: dns.RR_Header{
				Name:   question.Name,
				Rrtype: dns.TypeA,
				Class:  dns.ClassINET,
				Ttl:    0,
			},
			A: net.ParseIP("127.0.0.1").To4(),
		}
		resp.Answer = append(resp.Answer, &recordA)
	}
	err := writer.WriteMsg(&resp)
	if err != nil {
		return
	}
}
func main() {
	dns.HandleFunc(".", handler)
	err := dns.ListenAndServe(":53", "udp", nil)
	if err != nil {
		log.Println(err)
	}

}
