package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 	now := time.Now()

	// 	todayMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	// 	// Calculate the time difference
	// 	timeDiff := now.Unix() - todayMidnight.Unix()
	// 	fmt.Println(timeDiff)
	// 	fmt.Println(now.Unix())
	ls := 678.2399999999996
	re := fmt.Sprintf("%.4f", ls)
	o, _ := strconv.ParseFloat(re, 64)
	fmt.Println(o)
}
