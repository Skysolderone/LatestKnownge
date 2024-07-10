package main

import "fmt"

var i = 0

func main() {
	s := make(chan struct{}, 1)
	// s <- struct{}{}
	go func() {
		for i < 10 {
			<-s
			fmt.Println(i)
			i++
		}
	}()
	go func() {
		for i < 10 {
			s <- struct{}{}
			fmt.Println(i)
			i++
		}
	}()
	for {
	}
}
