package main

import (
	"fmt"
	"time"
)

func main() {
	front := time.Now()
	fmt.Println(front.Unix())
	time.Sleep(time.Second * 15)
	fmt.Println(time.Now().Unix() - front.Unix())
}
