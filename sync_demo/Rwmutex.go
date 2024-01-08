package main

import (
	"fmt"
	"sync"
)

var counter = 1
var rwmutex sync.RWMutex

func main() {
	go func() {
		rwmutex.RLock()
		defer rwmutex.RUnlock()
		fmt.Println("read:", counter)
	}()
	go func() {
		rwmutex.RLock()
		defer rwmutex.RUnlock()
		fmt.Println("read1:", counter)
	}()
	go func() {
		rwmutex.Lock()
		defer rwmutex.Unlock()
		counter++
	}()
	for {
	}
}
