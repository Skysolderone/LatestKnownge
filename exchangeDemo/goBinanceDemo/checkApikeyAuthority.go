package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adshao/go-binance/v2"
	"github.com/spf13/viper"
)

func main() {
	// /sapi/v1/account/apiRestrictions
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	apikey := viper.GetString("apikey")
	secretkey := viper.GetString("secretKey")
	client := binance.NewClient(apikey, secretkey)
	client.TimeOffset = -1000
	result, err := client.NewGetAPIKeyPermission().Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	err = client.NewCreateOrderService().
		Symbol("BTCUSDT").
		Side(binance.SideTypeBuy).
		Type(binance.OrderTypeMarket).
		Quantity("1").
		Test(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
