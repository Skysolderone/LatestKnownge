package main

import (
	"fmt"
	"sync"
)

func main() {
	var symap sync.Map

	symap.Store("1", "test1")
	symap.Store("2", "test2")
	symap.Store("3", "test3")
	fmt.Println(symap.Load("1"))
	symap.Delete("2")
	symap.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
