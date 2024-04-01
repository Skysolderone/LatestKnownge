package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type SyncJob struct {
	*sync.Cond
	holding   int32
	lastErr   error
	syncPoint *sync.Once
	syncFunc  func(interface{}) error
}

func NewSyncJob(fn func(interface{}) error) *SyncJob {
	return &SyncJob{
		syncFunc:  fn,
		syncPoint: &sync.Once{},
		Cond:      sync.NewCond(&sync.Mutex{}),
	}
}

func (s *SyncJob) Do(job interface{}) error {
	s.L.Lock()
	if s.holding > 0 {
		s.Wait()
	}
	s.holding += 1
	syncPoint := s.syncPoint
	s.L.Unlock()

	syncPoint.Do(func() {
		s.lastErr = s.syncFunc(job)
		s.L.Lock()
		fmt.Printf("%v\n", s.holding)
		s.holding = 0
		s.Broadcast()
		s.L.Unlock()
	})
	return s.lastErr
}

func main() {
	file, err := os.OpenFile("hello.txt", os.O_CREATE|os.O_RDWR, 0o700)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	syncjob := NewSyncJob(func(i interface{}) error {
		fmt.Println("do func ......\n")
		file.WriteString("COUNTER")
		time.Sleep(time.Second)
		return file.Sync()
	})
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("write...\n")
			syncjob.Do(file)
		}()
	}
	file.WriteString("hello \nit ")
	wg.Wait()
}
