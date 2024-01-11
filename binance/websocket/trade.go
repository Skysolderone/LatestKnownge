package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func main() {
	api_url := "wss://stream.binance.us:9443/ws/icpusdt@kline_1m"
	// resp, err := http.Get(api_url)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// fmt.Println(resp.Body)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	c, _, err := websocket.DefaultDialer.Dial(api_url, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	go func() {

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Println("MESSAGE:", string(message))
		}
	}()
	for {
		select {
		case <-interrupt:
			os.Exit(1)
		}
	}
}
