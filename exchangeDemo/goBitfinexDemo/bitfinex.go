package main

import (
	"log"
	"time"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/order"
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	apikey := viper.GetString("apikey")
	secretkey := viper.GetString("secretKey")
	c := rest.NewClient().Credentials(apikey, secretkey)
	// client := bitfinex.NewClient().Credentials(apikey, secretkey)

	// create order
	response, err := c.Orders.SubmitOrder(&order.NewRequest{
		Symbol: "BTCUSDT",
		CID:    time.Now().Unix() / 1000,
		Amount: 0.02,
		Type:   "EXCHANGE LIMIT",
		Price:  5000,
	})
	if err != nil {
		panic(err)
	}
	log.Println(response)
}
