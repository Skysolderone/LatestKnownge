package main

import (
	"fmt"
	"time"
)

func main3() {
	CountDown := int64(1730710200)
	t := time.Now().Unix()
	if t > CountDown && t < (CountDown+3600) {
		fmt.Println(3)
	}
}
