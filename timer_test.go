package test

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	timer := time.NewTimer(time.Second * 2)
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		ch <- "time*2"
	}()
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-timer.C:
		fmt.Println("time out")

	}
}
