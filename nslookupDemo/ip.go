package main

import (
	"fmt"
	"net"
)

func main() {
	//获取ip地址
	ips, err := net.LookupIP("www.google.com")
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
	//获取域名
	names, _ := net.LookupAddr("8.8.8.8")
	fmt.Println(names)
	//获取域名是否可用
	_, err = net.LookupNS("www.google.com")
	if err != nil {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}
