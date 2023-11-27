package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//basic
	pool := &sync.Pool{
		New: func() any {
			println("create pool")
			return struct{}{}
		},
	}
	buffer := pool.Get() //如果没有元素 则会调用new方法
	pool.Put(buffer)

	//level1
	pool1 := &sync.Pool{
		New: add, //因为 sync.Pool 只是本身的 Pool 数据结构是并发安全的，并不是说 Pool.New 函数一定是线程安全的。
		//Pool.New 函数可能会被并发调用 ，如果 New 函数里面的实现是非并发安全的，那就会有问题。
	}
	numWorkers := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			buffer2 := pool1.Get()
			_ = buffer2.(*[]byte)
			defer pool1.Put(buffer2)
		}()
	}
	wg.Wait()
	fmt.Println(createsum)
}
func add() any {
	atomic.AddInt32(&createsum, 1)
	buffer := make([]byte, 1024)
	return &buffer
}

var createsum int32
