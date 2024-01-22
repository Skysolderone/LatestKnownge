package main

import (
	"fmt"
	"net"
)

func main() {
	ips, err := net.LookupHost("www.google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
