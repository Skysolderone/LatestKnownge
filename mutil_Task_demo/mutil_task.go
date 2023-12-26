package main

import (
	"fmt"
	"sync"
)

type Order struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func main() {
	taskNum := 10
	orderCh := make(chan Order, taskNum)
	errCh := make(chan error, taskNum)
	var wg sync.WaitGroup
	for i := 0; i < taskNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//do something
		}()

	}
	orderList := make([]Order, taskNum)
	for i := 0; i < taskNum; i++ {
		select {
		case order, ok := <-orderCh:
			if ok {
				orderList = append(orderList, order)
			}
		case err := <-errCh:
			if err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("DONE")
		}
	}
	close(orderCh)
	close(errCh)
	wg.Wait()

}

// func processTask(task Task) {
// 	//do something
// }
