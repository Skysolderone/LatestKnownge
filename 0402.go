package main

import (
	"log"
	"math"
)

func main() {
	// s := [][]string{{"1", "2"}, {"3", "4"}}
	// for _, v := range s {
	// 	v[0] = "5"
	// 	log.Println(v[0])
	// }
	// log.Println(s)
	qty := 0.088193
	val := 4.00000
	rest := uint(FloorN(qty*(1.0/val), 0))
	log.Println(rest)
}

func FloorN(num float64, n int8) float64 {
	num += 0.00000000000001
	return math.Floor(num*math.Pow(10, float64(n))) / math.Pow(10, float64(n))
}
