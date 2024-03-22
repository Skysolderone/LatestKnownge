package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	num := -12345.52142
	absNum := math.Abs(float64(num))
	str := strconv.FormatFloat(absNum, 'f', -1, 64)
	fmt.Println(str)
}
