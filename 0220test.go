package main

import (
	"fmt"
	"sync"
	"time"
)

type MyMutex struct {
	ch chan bool
}

func NewMyMutex() *MyMutex {
	return &MyMutex{
		// Buffer size must be exactly one
		ch: make(chan bool, 1),
	}
}

// Lock locks m.
// If the lock is already in use, the calling goroutine
// blocks until the mutex is available.
func (m *MyMutex) Lock() {
	m.ch <- true
}

// Unlock unlocks m.
func (m *MyMutex) Unlock() {
	<-m.ch
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)

	bathroom := sync.Mutex{}

	takeAShower := func(name string) {
		defer wg.Done()

		fmt.Printf("%s: I want to take a shower. I'm trying to acquire the bathroom\n", name)
		bathroom.Lock()
		fmt.Printf("%s: I have the bathroom now, taking a shower\n", name)
		time.Sleep(500 * time.Microsecond)
		fmt.Printf("%s: I'm done, I'm unlocking the bathroom\n", name)
		bathroom.Unlock()
	}

	go takeAShower("Partier")
	go takeAShower("Candier")
	go takeAShower("Stringer")
	go takeAShower("Swimmer")

	wg.Wait()
	fmt.Println("main: Everyone is Done. Shutting down...")
}
