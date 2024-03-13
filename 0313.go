package main

import (
	"log"
	"sync"
)

func main() {
	var ms sync.Map
	var bs sync.Map
	ms.Store(10, 7.001)
	bs.Store(10, 5.002)
	max, _ := ms.Load(10)
	min, _ := bs.Load(10)
	ls, _ := max.(float64)
	rs, _ := min.(float64)
	res := ls - rs
	log.Println(res)
	if res > 1 {
		log.Println("success")
	}
	// log.Println(max)
	// log.Println(min)
	// if a != b && a-b > 1 {
	// 	log.Println("success")
	// }
}
