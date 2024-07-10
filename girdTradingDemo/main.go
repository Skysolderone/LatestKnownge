package main

import "fmt"

func main() {
	start := 1
	difference := 2
	count := 5

	// 生成等差数列
	arithmeticProgression := make([]int, 0, count)
	for i := 0; i < count; i++ {
		arithmeticProgression = append(arithmeticProgression, start+i*difference)
	}

	// 输出结果
	fmt.Println(arithmeticProgression) // 输出：[1 3 5 7 9]
}
