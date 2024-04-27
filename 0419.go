package main

import (
	"log"
	"strings"
)

func main() {
	s := "abccd123"
	symbol := strings.Trim(s, "abc")
	log.Println(symbol)
}
