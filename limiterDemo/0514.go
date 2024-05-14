package main

import (
	"context"
	"fmt"

	"golang.org/x/time/rate"
)

func main() {
	ls := rate.NewLimiter(100, 100)
	for i := range 1000000 {
		ls.Wait(context.Background())

		fmt.Println(i)
	}
	// limiter := rate.NewLimiter(10, 100)

	// // 模拟处理1000个请求
	// for i := 0; i < 1000; i++ {
	// 	// 等待直到可以获得一个令牌
	// 	if err := limiter.Wait(context.Background()); err != nil {
	// 		fmt.Println("Error:", err)
	// 		continue
	// 	}
	// 	// 模拟处理请求
	// 	fmt.Println("Handling request", i+1)
	// }
}
