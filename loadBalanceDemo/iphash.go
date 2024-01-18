package main

import (
	"fmt"
	"hash/fnv"
)

type IPhash struct {
	servers []string
}

func NewIPHash(servers []string) *IPhash {
	return &IPhash{
		servers: servers,
	}
}
func (ih *IPhash) GetServerByIp(ip string) string {
	h := fnv.New32a()
	h.Write([]byte(ip))
	index := int(h.Sum32()) % len(ih.servers)
	return ih.servers[index]
}

// ip hash
func main() {
	servers := []string{"String1", "String2", "String3"}
	ih := NewIPHash(servers)
	ips := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3"}
	for _, ip := range ips {
		fmt.Printf("Request from IP %s sent to:%s\n", ip, ih.GetServerByIp(ip))
	}
}
