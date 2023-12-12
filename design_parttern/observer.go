package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Subject struct {
	observers []chan string
	mu        sync.Mutex
}

func (s *Subject) AddObserver(obs chan string) {
	s.mu.Lock()
	s.observers = append(s.observers, obs)
	s.mu.Unlock()
}
func (s *Subject) RemoveObserver(obs chan string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, observer := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}
func (s *Subject) NotifyObserver(msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, observer := range s.observers {
		go func(ch chan string) {
			ch <- msg
		}(observer)
	}
}
func main() {
	subject := &Subject{}
	obs1 := make(chan string)
	subject.AddObserver(obs1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for message := range obs1 {
			fmt.Println("obs1 receive msg:", message)
		}
		wg.Done()
	}()
	obs2 := make(chan string)
	subject.AddObserver(obs2)
	go func() {
		for message := range obs2 {
			fmt.Println("obs2 receive msg:", message)
		}
		wg.Done()
	}()
	subject.NotifyObserver("Hello observers")
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	go func() {
		<-ch
		close(obs1)
		close(obs2)
	}()

	wg.Wait()
}
