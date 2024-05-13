package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// viper.SetConfigName("config.yaml")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }
	// ls := viper.Get("where")
	// fmt.Println(ls)
	// str := viper.Get("hello")
	// fmt.Println(str)
	viper.SetConfigName("cfg.toml")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	str := viper.GetString("540008link.binance")
	fmt.Println(str)
}
