package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
)

//config 存储原本币种文件
func main() {
	file1, err := os.Create("./log.log")
	if err != nil {
		panic("create err")
	}
	log.SetOutput(file1)

	// log.SetFlags(log.LstdFlags)
	// body, err := os.ReadFile("./config")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(body))
	file, err := os.Open("./config")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	res := make(map[string]int)

	client := binance.NewClient("", "")
	client.BaseURL = "http://api.binance.us"
	spotExchangeInfo, err := client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		panic(err)
	}
	// obj := make(map[string]int, 0)
	for _, v := range spotExchangeInfo.Symbols {
		// log.Println(v.Symbol[:len(v.Symbol)-4])
		res[v.Symbol[:len(v.Symbol)-4]] = 1

	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		if res[scanner.Text()] != 1 {
			continue
		}
		log.Println(scanner.Text())
	}
	// fmt.Println(obj)
}
