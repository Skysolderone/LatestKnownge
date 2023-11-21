package main

import (
	"fmt"
	"runtime"
)

// 向map添加 n 个元素，然后删除所有元素意味着在内存中保持相同数量的存储桶。
// 因此，我们必须记住，由于 Go map只能增长，因此其内存消耗也会随之增加。它没有自动化的策略来缩小。如
// 果这导致内存消耗过高，我们可以尝试不同的选项，比如强制 Go 重新创建map或使用指针来检查是否可以进行优化。
func main() {
	n := 1_000_000
	m := make(map[int]*[128]byte)
	printAlloc()
	for i := 0; i < n; i++ {
		m[i] = &[128]byte{}
	}
	printAlloc()
	for i := 0; i < n; i++ {
		delete(m, i)
	}
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}
