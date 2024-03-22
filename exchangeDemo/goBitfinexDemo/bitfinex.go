package main

import (
	"fmt"
	"log"
	"strings"
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
	apikey = strings.TrimSpace(apikey)
	secretkey = strings.TrimSpace(secretkey)
	log.Println(apikey)
	log.Println(secretkey)
	// key := "23d2808deb79a7cd31edf4f00361ccabf9b85bb471b"
	// secret := "0a1baa74959a961463406c786459c5562bbaedaa92b"
	client := rest.NewClientWithURL("https://test.bitfinex.com/v2/").Credentials(apikey, secretkey)

	// wallets, err := client.Wallet.Wallet()
	// if err != nil {
	// 	log.Fatalf("getWallets %s", err)
	// }
	// fmt.Println(wallets)
	response, err := client.Orders.SubmitOrder(&order.NewRequest{
		Symbol: "tBTCUSD",
		CID:    time.Now().Unix() / 1000,
		Amount: 0.02,
		Type:   "EXCHANGE LIMIT",
		Price:  5000,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
