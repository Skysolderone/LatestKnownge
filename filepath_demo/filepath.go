package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// 目录遍历
// basic
func visit(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}

func main() {
	root := "../."
	//basic
	// err := filepath.Walk(root, visit)
	// if err != nil {
	// 	fmt.Printf("%v:%v\n", root, err)
	// 	return
	// }

	//signal thread
	// sequentialTraversal(root)

	//mutil thread
	ch := make(chan struct{}, 10)
	var wg sync.WaitGroup
	// 获取CPU核心数
	//资源优化
	numCPU := runtime.NumCPU()

	// 设置GOMAXPROCS为CPU核心数
	runtime.GOMAXPROCS(numCPU)

	// 根据CPU核心数划分工作
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go concurrentTraversal(root, &wg, ch)
	}

	// wg.Add(1)
	// go concurrentTraversal(root, &wg, ch)
	go func() {
		wg.Wait()
		close(ch)
	}()
	select {
	case <-ch:
		// 遍历完成，没有错误
		return
	default:
		// 有错误发生，进行处理
		log.Fatal("Concurrent traversal failed")
	}
}

var mu sync.Mutex

// signale thread process
func processFile(file string) {
	mu.Lock()
	defer mu.Unlock()
	//process
	fmt.Println("file:", file)
}
func sequentialTraversal(root string) {
	files, err := os.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			sequentialTraversal(filepath.Join(root, file.Name()))
		} else {
			processFile(filepath.Join(root, file.Name()))
		}
	}

}

// mutil thread process
func concurrentTraversal(root string, wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	files, err := os.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			go concurrentTraversal(filepath.Join(root, file.Name()), wg, ch)
		} else {
			processFile(filepath.Join(root, file.Name()))
		}

	}
	ch <- struct{}{}
}
