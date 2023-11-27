package main

import (
	"fmt"
)

// 使用通道顺序输出
var s = make(chan int)
var s1 = make(chan int)

func main() {

	go one()
	go two()

	//time.Sleep(time.Second * 3)
	for {
	}
}

func one() {
	for i := 1; i < 10; i++ {
		s1 <- 2*i - 1
		result := <-s
		fmt.Println(result)
		//		return

	}
}

func two() {
	for i := 1; i < 10; i++ {
		result := <-s1
		fmt.Println(result)
		s <- 2 * i
		//		return
	}
}
