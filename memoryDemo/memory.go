package main

import (
	"fmt"
	"runtime"
)

func main() {

	//basic
	//获取当前分配内存的大小
	memStats := new(runtime.MemStats)
	runtime.ReadMemStats(memStats)
	fmt.Println("heapalloc:", memStats.HeapAlloc)
	//手动分配内存
	data := make([]byte, 1024)
	_ = data
	//再次获取内存状态
	runtime.ReadMemStats(memStats)
	fmt.Println("heapalloc:", memStats.HeapAlloc)

	//gc
	obj := new(Object)
	_ = obj
	runtime.GC()
	fmt.Println("GC COMPLETE")
}

type Object struct {
	data int
}
