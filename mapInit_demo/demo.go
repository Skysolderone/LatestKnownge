package main

import (
	"fmt"
	"math"
	"sort"
)

var demo = make(map[float64]float64, 0)

func main() {
	demo[3.24] += 2.3
	fmt.Println(demo)
	data := []int{1, 3, 5, 2, 10}
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	fmt.Println(data)
	d := 84819.4000
	fmt.Println(d / 10.0)
	pri := math.Round(d/10.0) * 10
	fmt.Println(pri)
}
