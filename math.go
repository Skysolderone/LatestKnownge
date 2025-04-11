package main

import (
	"fmt"
	"math"
)

func main() {
	rate := (100.00 / 300.00) * 100
	fmt.Println(rate)
	rateresult := math.Floor(rate)
	fmt.Println(rateresult)
}
