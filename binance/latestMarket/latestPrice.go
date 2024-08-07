package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	payload := `
{
    "id": "9d32157c-a556-4d27-9866-66760a174b57",
    "method": "ticker.price",
    "params": {
        "symbol": "BTCUSDT"
    }
}	
`
	url := "wss://stream.binance.com:9443/ws"
	client, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	client.WriteMessage(websocket.TextMessage, []byte(payload))
	for {
		_, s, err := client.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(s))

	}
}
