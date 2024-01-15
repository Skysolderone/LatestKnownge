package main

import (
	"fmt"
	"math"
)

//质数检测器

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	k := math.Sqrt(float64(num))
	for i := 3; i <= int(k); i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var num int
	fmt.Scanln(&num)
	fmt.Println(isPrime(num))
}
