package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		panic("empty slice")
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	intSlice := []int{3, 1, 4, 2}
	floatSlcie := []float64{3.14, 1.61, 4.67, 2.71}
	fmt.Println(Min[int](intSlice))
	fmt.Println(Min[float64](floatSlcie))
}
