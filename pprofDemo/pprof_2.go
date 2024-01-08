// main.go
package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)
//http://localhost:6060/debug/pprof/
func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// 模拟内存泄漏
	for {
		data := make([]byte, 1024)
		_ = data
		time.Sleep(time.Millisecond * 10)
	}
}
