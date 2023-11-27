package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var cond = sync.NewCond(&mutex)
var queue []int

func producer() {
	i := 0
	for {
		cond.L.Lock()
		queue = append(queue, i)
		i++
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(1 * time.Second)
	}
}
func consumer(str string) {
	for {
		cond.L.Lock()
		for len(queue) == 0 {
			cond.Wait()
		}
		fmt.Println(str, queue[0])
		queue = queue[1:]
		cond.L.Unlock()

	}
}

func main() {
	go producer()

	go consumer("comsumer-1")
	go consumer("comsumer-2")
	for {
		time.Sleep(time.Second * 1)
	}

}
