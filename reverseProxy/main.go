package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net"
	"strings"
)

func GetMacADDRESS() string {
	netInterface, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	var macaddress []string
	for _, inter := range netInterface {
		macAddr := inter.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macaddress = append(macaddress, macAddr)
	}
	str := strings.Join(macaddress, "_")
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
