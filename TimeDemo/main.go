package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix() - 1729828076)
	if time.Now().Unix()-1729828076 > 604800 {
		fmt.Println(true)
	}
}
