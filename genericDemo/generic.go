package main

import "fmt"

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	intSlice := []int{1, 2, 3}
	stringSlice := []string{"hello", "world"}
	PrintSlice[int](intSlice)
	PrintSlice[string](stringSlice)
}
