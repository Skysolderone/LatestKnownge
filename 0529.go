package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "MATIC-USDT"
	s := strings.Replace(str, "-", "", -1)
	if !strings.Contains(s, "USDT") {
		s = s + "USDT"
	}
	fmt.Println(s)
}
