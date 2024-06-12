package main

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("spot.yaml")
	viper.ReadInConfig()
	spotSymbol := viper.GetStringMapString("spot")
	file, _ := os.OpenFile("final.yaml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o755)
	defer file.Close()
	viper.SetConfigFile("futures.yaml")
	viper.ReadInConfig()
	futuresSymbol := viper.GetStringMapString("spot")
	for i := range spotSymbol {
		symbol := i
		if _, ok := futuresSymbol[symbol]; ok {
			file.WriteString(strings.ToUpper(symbol) + ":" + "1\n")
		}

	}
}
