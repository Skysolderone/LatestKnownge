package main

import "fmt"

func main() {
	ls := make([]int, 0)

	for i := 1; i < 6; i++ {
		ls = append(ls, i)
	}
	for _, v := range ls {
		go func(int) {
			fmt.Println("%t", &v)
		}(v)
	}
	for {
	}
}
