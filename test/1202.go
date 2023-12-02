package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var x int32
	var y int32
	var z int32

	go func() {
		x = atomic.AddInt32(&x, 1)
		y++
		z = y + x
	}()
	time.Sleep(time.Second)
	fmt.Println(x, y, z)
}
