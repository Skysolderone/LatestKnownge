package main

import (
	"fmt"
	"log"

	"github.com/bitmartexchange/bitmart-go-sdk-api"
)

type Spot struct {
	client *bitmart.CloudClient
}

func NewSpot(AccessKey, SecretKey, endpoint, Memo string) *Spot {
	config := bitmart.Config{}

	config.ApiKey = AccessKey
	config.SecretKey = SecretKey
	config.Url = endpoint
	config.Memo = Memo
	// fmt.Println(endpoint)
	config.Headers = map[string]string{
		"x-ex": "bitmart_spot",
	}

	client := bitmart.NewClient(config)
	// fmt.Println(client.Config)
	return &Spot{
		client: client,
	}
}

func (b *Spot) GetOrder(id string) {
	order, err := b.client.GetSpotOrderByOrderId(id, "history", 10000)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(order)
}

func (b *Spot) Balance() {
	accounts, _ := b.client.GetSpotAccountWallet()
	// fmt.Printf("%#v", accounts)

	fmt.Println(accounts)
}

func main() {
	api := "39237ed9824090889dfc875e39d41a47dc56bf49"
	sec := "83fa4361a186a7e3e074c8dd9c755dd51ac016e1ff86ff628b01ba6e25695970"
	ProxyURL := "https://api-cloud.bitmart.com"
	Memo := "uTrading"
	s := NewSpot(api, sec, ProxyURL, Memo)
	s.Balance()
	s.GetOrder("896046636051021312")
}
