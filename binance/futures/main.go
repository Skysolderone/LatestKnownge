package main

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2/futures"
)

func main() {
	client := futures.NewClient("", "")
	client.BaseURL = ""
	Info, err := client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(Info.Symbols)
}
