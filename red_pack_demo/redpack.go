package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type RedPack struct {
	Less     float64 //剩余金额
	LessPack int     //剩余个数

}

func remainTwoDecimal(num float64) float64 {
	numStr := strconv.FormatFloat(num, 'f', 2, 64)
	num, _ = strconv.ParseFloat(numStr, 64)
	return num
}
func getRandomRedPack(rp *RedPack) float64 {
	if rp.LessPack < 0 {
		return 0
	}
	if rp.LessPack == 1 {
		return remainTwoDecimal(rp.Less + 0.01)
	}
	avgAmount := math.Floor(100*(rp.Less/float64(rp.LessPack))) / float64(100)
	avgAmount = remainTwoDecimal(avgAmount)
	rand.NewSource(time.Now().UnixNano())
	var max float64
	if avgAmount > 0 {
		max = 2*avgAmount - 0.01
	} else {
		max = 0
	}
	money := remainTwoDecimal(rand.Float64()*(max) + 0.01)
	rp.LessPack -= 1
	rp.Less = remainTwoDecimal(rp.Less + 0.01 - money)
	return money
}
func main() {
	red := &RedPack{
		Less:     0.16,
		LessPack: 5,
	}
	red.Less -= 0.01 * float64(red.LessPack)
	total := red.LessPack
	for i := 0; i < total; i++ {
		fmt.Println(getRandomRedPack(red))
	}
}
