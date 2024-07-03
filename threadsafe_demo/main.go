package main

import (
	"fmt"

	"github.com/hayageek/threadsafe"
)

// 性能会比原生低  不建议使用
func main() {
	arr := threadsafe.NewArray[int](5)

	for i := 0; i < arr.Length(); i++ {
		arr.Set(i, i*10)
	}
	for i := 0; i < arr.Length(); i++ {
		value, _ := arr.Get(i)
		fmt.Println(value)
	}
}
