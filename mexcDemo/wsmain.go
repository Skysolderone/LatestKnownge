package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
)

const url = "wss://wbs.mexc.com/ws"

func main() {
	header := http.Header{}
	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	channel := `{
    "method": "SUBSCRIPTION",
    "params": [
        "spot@public.miniTickers.v3.api@UTC+8"
    ]
}`
	err = conn.WriteMessage(websocket.TextMessage, []byte(channel))
	if err != nil {
		log.Fatal("WebSocket write error:", err)
	}
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)

			return
		}
		res := wsData{}
		sonic.Unmarshal(message, &res)
		for _, v := range res.Tickers {
			fmt.Println(v.Symbol)
		}
	}
}

type Ticker struct {
	Symbol        string `json:"s"`
	Price         string `json:"p"`
	Rate          string `json:"r"`
	TradeRate     string `json:"tr"`
	High          string `json:"h"`
	Low           string `json:"l"`
	Volume        string `json:"v"`
	QuoteVolume   string `json:"q"`
	LastRT        string `json:"lastRT"`
	MarketType    string `json:"MT"`
	NotionalValue string `json:"NV"`
}

type wsData struct {
	Tickers []Ticker `json:"d"`
	Channel string   `json:"c"`
	Time    int64    `json:"t"`
}
