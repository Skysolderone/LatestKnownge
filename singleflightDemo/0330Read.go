package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	var g singleflight.Group
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			v, err, shared := g.Do("objectOk", func() (any, error) {
				fmt.Println("GOROUTINE %v RUNNING..\n", idx)
				time.Sleep(2 * time.Second)
				return "objectvalue", nil
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("goroutine %v result %v shared %v\n", idx, v, shared)
		}(i)
	}
	wg.Wait()
}
