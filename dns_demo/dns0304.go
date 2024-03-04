package main

import (
	"fmt"
	"log"
	"time"

	"github.com/miekg/dns"
)

func resolve(domain string, qtype uint16) []dns.RR {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), qtype)
	m.RecursionDesired = true
	c := &dns.Client{Timeout: 5 * time.Second}
	response, _, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		log.Fatal(err)
	}
	if response == nil {
		log.Fatal("no response")
	}
	for _, answer := range response.Answer {
		fmt.Printf("%s\n", answer.String())
	}
	return response.Answer
}

func main() {
	domain := "www.baidu.com"
	resolve(domain, dns.TypeA)
}
