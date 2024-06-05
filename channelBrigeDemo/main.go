package main

import (
	"fmt"
	"sync"
)

func orDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case val, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

// bridge converts a channel of channels into a single channel.
func bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	var wg sync.WaitGroup

	go func() {
		defer close(valStream)

		for {
			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if !ok {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}

			wg.Add(1)
			go func(s <-chan interface{}) {
				defer wg.Done()
				for val := range orDone(done, s) {
					select {
					case valStream <- val:
					case <-done:
						return
					}
				}
			}(stream)
		}
	}()

	go func() {
		wg.Wait()
		close(valStream)
	}()

	return valStream
}

// genVals creates a channel of channels with values.
func genVals() <-chan <-chan interface{} {
	chanStream := make(chan (<-chan interface{}))
	go func() {
		defer close(chanStream)
		for i := 0; i < 10; i++ {
			stream := make(chan interface{}, 1)
			stream <- i
			close(stream)
			chanStream <- stream
		}
	}()
	return chanStream
}

func main() {
	done := make(chan interface{})
	defer close(done)
	// fmt.Println("TEST")
	for v := range bridge(done, genVals()) {
		fmt.Printf("%v ", v)
	}
	select {}
}
