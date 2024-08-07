package main

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func main() {
	bookTickerHander := func(event *binance.WsDepthEvent) {
		fmt.Printf("%#v", event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsDepthServe("BTCUSDT", bookTickerHander, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	<-doneC
}
