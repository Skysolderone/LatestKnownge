package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "XBTUSDT"
	s = strings.Replace(s, "BTC", "XBT", 1) + "M"
	fmt.Println(s)
}
