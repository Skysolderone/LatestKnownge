package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁
var counter int
var mutex sync.Mutex

// 读写锁
var data map[string]string
var rwMutex sync.RWMutex

func main() {
	//互斥锁
	var wg sync.WaitGroup
	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)
	// 	go incomer(i, &wg)
	// }
	// wg.Wait()
	// fmt.Printf("Final counter :%d \n", counter)

	//读写锁 读多写少的场景中提高并发性能
	data = make(map[string]string)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Read(i, &wg)
	}
	//for i := 0; i < 3; i++ {//模拟写锁饥饿//
	/*写锁饥饿是指当有读锁持有时，写锁一直无法获取的情况。

	要避免写锁饥饿，应该尽量减小读操作的临界区，避免长时间占用读锁。*/
	wg.Add(1)
	go Write(1, &wg)
	//}
	wg.Wait()

}
func Read(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.RLock()
	fmt.Printf("goroutine %d:read\n", i)
	time.Sleep(500 * time.Millisecond)
	rwMutex.RUnlock()
}
func Write(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.Lock()
	key := fmt.Sprintf("key%d", i)
	value := fmt.Sprintf("value%d", i+10)
	data[key] = value
	fmt.Printf("goroutine %d:write\n", i)
	time.Sleep(500 * time.Millisecond)
	rwMutex.Unlock()
}
func incomer(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		mutex.Lock()
		counter++
		fmt.Printf("goroutine :%d worker\n", i)
		mutex.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}
