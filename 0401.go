package main

import (
	"fmt"
	"strings"
)

func main() {
	change := "5000.0,sell,83"
	price := strings.Split(change, ",")
	fmt.Println(price[0])
}
