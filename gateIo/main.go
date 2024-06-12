package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gateio/gateapi-go/v6"
	"github.com/spf13/viper"
)

func main() {
	config := gateapi.NewConfiguration()
	config.Debug = true
	config.Key = "7d949dac241b26c2fda790c5e8d3d2f2"
	config.Secret = "a51bac99177efce818907b391845a76ff0a717e200ee08a38da5a479fe1fa339"
	client := gateapi.NewAPIClient(config)
	// client.ChangeBasePath()
	// uncomment the next line if your are testing against testnet
	// client.ChangeBasePath("https://fx-api-testnet.gateio.ws/api/v4")
	// ctx := context.WithValue(context.Background(),
	// 	gateapi.ContextGateAPIV4,
	// 	gateapi.GateAPIV4{
	// 		Key:    "7d949dac241b26c2fda790c5e8d3d2f2",
	// 		Secret: "a51bac99177efce818907b391845a76ff0a717e200ee08a38da5a479fe1fa339",
	// 	},
	// )
	// viper.SetConfigFile("cfg.yaml")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	client.ChangeBasePath("https://api.gateio.ws/api/v4")
	symbols := viper.GetStringMapString("spot")
	fmt.Println(len(symbols))
	file, err := os.OpenFile("spot.yaml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o755)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("spot:\n")
	defer file.Close()
	result, _, err := client.FuturesApi.GetFuturesContract(context.Background(), "usdt", "BTC_USDT")
	if err != nil {
		if e, ok := err.(gateapi.GateAPIError); ok {
			fmt.Printf("gate api error: %s\n", e.Error())
		} else {
			fmt.Printf("generic error: %s\n", err.Error())
		}
	}
	fmt.Println(result)
	// for _, v := range result {
	// 	// fmt.Println(v.Quote)
	// 	if v.Quote != "USDT" {
	// 		continue
	// 	}
	// 	symbol := v.Base
	// 	if _, ok := symbols[strings.ToLower(symbol)]; ok {
	// 		_, err := file.WriteString("\t\t" + symbol + ":" + " 0\n")
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}

	// }
	// fmt.Println(-1 - (-2))
}
