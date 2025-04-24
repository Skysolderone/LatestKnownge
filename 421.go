package main

import (
	"fmt"
	"math"
)

func main() {
	s := (0 / 3) * 100.00
	floorN(s, 2)
	fmt.Println(s)
}

func floorN(num float64, n int8) float64 {
	num += 0.00000000000001
	return math.Floor(num*math.Pow(10, float64(n))) / math.Pow(10, float64(n))
}
