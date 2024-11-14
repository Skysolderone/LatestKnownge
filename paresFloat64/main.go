package main

import (
	"fmt"
	"math"
)

func main() {
	s := 2257.7513214
	// multiplier := math.Pow(10, float64(4))
	fmt.Println(roundToDecimals(s, 4))
}

func roundToDecimals(value float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(value*multiplier) / multiplier
}
