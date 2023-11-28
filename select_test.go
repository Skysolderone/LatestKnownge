package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSlect(t *testing.T) {
	ch := make(chan string)
	ch1 := make(chan string)
	go func() {
		time.After(time.Second * 2)
		ch <- "goroutine1"

	}()
	go func() {
		time.After(time.Second * 1)
		ch1 <- "goroutine2"
	}()
	for {
		select {
		case res := <-ch:
			fmt.Println(res)
		case res1 := <-ch1:
			fmt.Println(res1)
		case <-time.After(time.Second * 3):
			fmt.Println("finish")
			return
		}
	}
}
