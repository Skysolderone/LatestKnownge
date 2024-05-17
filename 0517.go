package main

import (
	"fmt"
	"strings"
)

func main() {
	address := "Arbitrum One is not address"
	network := strings.Replace(strings.TrimLeft(address, "/register/"), "is not address", "", -1)
	fmt.Println(network)
}
