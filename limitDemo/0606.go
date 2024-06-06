package main

import (
	"fmt"

	"golang.org/x/time/rate"
)

func main() {
	limite := rate.NewLimiter(50/1, 1)
	i := 0
	for {
		i++
		if limite.Allow() {
			fmt.Println(i)
		}
	}
}
