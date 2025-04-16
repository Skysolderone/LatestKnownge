package main

import (
	"fmt"
	"time"
)

type Event struct {
	Name string
	Data interface{}
}

func main() {
	eventChan := make(chan Event, 10)
	go func() {
		for event := range eventChan {
			fmt.Printf("[Consumer] 处理事件: %s，数据: %v\n", event.Name, event.Data)
		}
	}()
	eventChan <- Event{"UserRegistered", "张三"}
	time.Sleep(time.Second)
}
