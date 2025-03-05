package main

import "fmt"

var (
	GlobalCheckBuy   map[int64]float64
	GlobalCheckSell  map[int64]float64
	GlobalCheckBuyc  map[int64]float64
	GlobalCheckSellc map[int64]float64
)

func InitGlobalEvent() {
	GlobalCheckBuy = make(map[int64]float64, 100/50)
	GlobalCheckSell = make(map[int64]float64, 100/50)
	GlobalCheckBuyc = make(map[int64]float64, 100/50)
	GlobalCheckSellc = make(map[int64]float64, 100/50)
	for i := 0; i < 100; i += 50 {
		GlobalCheckBuy[int64(i)+50] = 0
		GlobalCheckSell[int64(i)+50] = 0
		GlobalCheckBuyc[int64(i)+50] = 0
		GlobalCheckSellc[int64(i)+50] = 0
	}
}

func main() {
	InitGlobalEvent()
	GlobalCheckBuy[50] = 5
	fmt.Println(GlobalCheckBuy)
	GlobalCheckBuy = GlobalCheckBuyc
	fmt.Println(GlobalCheckBuyc)
}
