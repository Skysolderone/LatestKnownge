package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
)

func main() {
	client := gateapi.NewAPIClient(gateapi.NewConfiguration())
	client.GetConfig().Debug = true
	ls := &gateapi.ListCandlesticksOpts{}
	limit := optional.NewInt32(1)
	ls.Interval = optional.NewString("1d")
	ls.Limit = limit
	// ls.To = optional.NewInt64(1718163744)
	resp, _, err := client.SpotApi.ListCandlesticks(
		context.Background(),
		"BTC_USDT",
		ls,
	)
	// params := &gateapi.ListFuturesCandlesticksOpts{}
	// params.Limit = optional.NewInt32(2)
	// params.To = optional.NewInt64(1718163744)
	// resp, _, err := client.FuturesApi.ListFuturesCandlesticks(
	// 	context.Background(),
	// 	"usdt",
	// 	"BTC_USDT",
	// 	params,
	// )
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range resp {
		change, _ := strconv.Atoi(v[0])

		fmt.Println(strconv.Itoa(change + 60 - 1))
	}
}
