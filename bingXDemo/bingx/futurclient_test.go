package main

import (
	"fmt"
	"testing"

	"bingx/bingx/market"
)

func TestGetSymboFu(t *testing.T) {
	// test get symbol detail
	client := new(market.MarketFutureClient).Init()
	parmas := make(map[string]string, 0)
	parmas["symbol"] = "BTC-USDT"
	resp, err := client.GetSymbolDetail(parmas)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(resp.Data.Symbols)
	for _, v := range resp.Data {
		fmt.Println(v.ApiStateOpen)
		if v.ApiStateOpen == "true" {
			fmt.Println("is bool")
		}
	}
}
