package main

import "fmt"

func main2() {
	var nowOfWeek uint = 3
	ts := []uint{1, 2}
	//[1,2]
	var s uint = 0
	var next uint = 0
	for _, v := range ts {
		if (nowOfWeek - v) > s {

			s = nowOfWeek - v

			next = v
		}
	}
	fmt.Println(next)
}
