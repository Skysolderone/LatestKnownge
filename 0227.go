package main

import "time"

var m = make(map[string]string)

func main() {
	go func() {
		for {
			_ = m["x"]
		}
	}()
	go func() {
		for {
			m["y"] = "foo"
		}
	}()
	time.Sleep(1 * time.Second)
}
