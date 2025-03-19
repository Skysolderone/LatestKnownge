package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "84210000213.4214000000000"
	s = strings.TrimRight(s, "0")
	fmt.Println(s)
}
