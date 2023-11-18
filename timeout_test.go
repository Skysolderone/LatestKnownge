package test

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// basic
func TestBasic(t *testing.T) {
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "hello"
	}()
	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(2 * time.Second):
		fmt.Println("time out")
	}
}

// http
func fetch(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	time.Sleep(time.Second)
	ch <- "request success"

}

func TestHttp(t *testing.T) {

	ch := make(chan string)
	url := "http://www.google.com"
	go fetch(url, ch)
	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second * 2):
		fmt.Println("request time out")
	}
}

//multi channel

func TestMultiChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	// 启动goroutine模拟耗时操作
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "Operation 1 completed!"
	}()

	// 启动goroutine模拟另一个耗时操作
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "Operation 2 completed!"
	}()

	// 使用select和time.After实现超时机制和多通道选择
	select {
	case result := <-ch1:
		fmt.Println(result)
	case result := <-ch2:
		fmt.Println(result)
	case <-time.After(4 * time.Second):
		fmt.Println("Timeout! Operations took too long.")
	}
}
