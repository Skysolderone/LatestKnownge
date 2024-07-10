package main

import (
	"fmt"
	"sync"
)

var S = map[string]string{
	"a": "1",
	"b": "2",
	"c": "3",
}

func main() {
	var Rm sync.RWMutex

	go func() {
		Rm.RLock()
		y, _ := S["a"]
		fmt.Println(y)
	}()
	go func() {
		Rm.RLock()
		y, _ := S["b"]
		fmt.Println(y)
	}()
	fmt.Println(S["c"])
	
	select {}
}
