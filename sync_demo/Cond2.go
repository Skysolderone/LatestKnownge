package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	condtionMutex sync.Mutex
	condition     *sync.Cond
	isReady       bool
)

func waitForCondition() {
	condtionMutex.Lock()
	defer condtionMutex.Unlock()
	for !isReady {
		fmt.Println("Waitting for condition")
		condition.Wait()
	}
	fmt.Println("Condition is pending")
}
func setCondition() {
	time.Sleep(2 * time.Second)
	condtionMutex.Lock()
	isReady = true
	condition.Signal()
	condtionMutex.Unlock()
}
func main() {
	condition = sync.NewCond(&condtionMutex)
	go waitForCondition()
	go setCondition()
	time.Sleep(time.Second * 5)
}
