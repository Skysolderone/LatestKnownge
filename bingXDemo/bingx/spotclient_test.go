package main

import (
	"fmt"
	"testing"

	"bingx/bingx/account"
	"bingx/bingx/market"
	"bingx/bingx/trade"
)

func TestGetSymSpot(t *testing.T) {
	// test get symbol detail
	client := new(market.MarketSpotClient).Init()
	parmas := make(map[string]string, 0)
	parmas["symbol"] = "BTC-USDT"
	resp, err := client.GetSymbolDetail(parmas)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(resp.Data.Symbols)
	for _, v := range resp.Data.Symbols {
		fmt.Println(v)
	}
}

func TestGetAccount(t *testing.T) {
	client := new(account.AccountSpotClient).Init()
	parmas := make(map[string]string, 0)
	// parmas["symbol"] = "BTC-USDT"
	resp, err := client.GetBanlanceDetail(parmas)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	// fmt.Println(resp.Data.Symbols)
	for _, v := range resp.Data.Balances {
		fmt.Println(v)
	}
}

func TestPlaceOrder(t *testing.T) {
	// 	限價單必須傳price參數。
	// 限價單必須傳quantity或quoteOrderQty其中一個，當两個參數同時傳遞時，服務端優先使用參數quantity。
	// 市價買單必須傳quoteOrderQty參數。
	// 市價賣單必須傳quantity參數。

	client := new(trade.TradeSpotClient).Init()
	parmas := make(map[string]string, 0)
	parmas["symbol"] = "BTC-USDT"
	resp, err := client.PlaceOrder(parmas)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(resp.Data.Symbols)
	fmt.Println(resp)
}

func TestGetOrder(t *testing.T) {
	client := new(trade.TradeSpotClient).Init()
	parmas := make(map[string]string, 0)
	parmas["symbol"] = "BTC-USDT"
	resp, err := client.GetOrder(parmas)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(resp.Data.Symbols)
	fmt.Println(resp)
}
