package main

import (
	"fmt"
	"log"
)

func main() {
	client := new(v1.MixMarketClient).Init()
	params := internal.NewParams()
	params["symbol"] = "BTCUSDT_UMCBL"
	// params["umcbl"] = "USDT"
	result, err := client.MarkPrice(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
