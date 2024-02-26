package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	apikey := viper.GetString("apikey")
	secretkey := viper.GetString("secretKey")
	fmt.Println(apikey)
	fmt.Println(secretkey)
}
