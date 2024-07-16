package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
)

func main() {
	wsurl := "wss://ws.okx.com:8443/ws/v5/public"
	msg := OkxRequest{}
	msg.Op = "subscribe"
	args := Args{}
	args.Channel = "liquidation-orders"
	args.InstType = "SWAP"
	msg.Args = append(msg.Args, args)
	ws, _, err := websocket.DefaultDialer.Dial(wsurl, nil)
	if err != nil {
		log.Println(err)
	}
	
	go func() {
		ticker := time.NewTicker(time.Second * 30)
		for range ticker.C {
			ws.WriteMessage(websocket.PingMessage, []byte("ping"))
		}
	}()
	sub, _ := sonic.Marshal(msg)
	ws.WriteMessage(websocket.TextMessage, sub)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			continue
		}
		data := OkxResponse{}
		sonic.Unmarshal(msg, &data)
		if len(data.Data) == 0 {
			continue
		}
		text := fmt.Sprintf(`
<b>OKX</b>
</b>交易对:%s</b>
成交额:%s
<b>方向:%s</b>
		`, data.Data[0].InstFamily, data.Data[0].Details[0].Sz, data.Data[0].Details[0].PosSide)
		fmt.Println(text)
	}
}

type OkxRequest struct {
	Op   string `json:"op"`
	Args []Args `json:"args"`
}
type Args struct {
	Channel  string `json:"channel"`
	InstType string `json:"instType"`
}

type OkxResponse struct {
	Arg  Arg    `json:"arg"`
	Data []Data `json:"data"`
}
type Arg struct {
	Channel  string `json:"channel"`
	InstType string `json:"instType"`
}
type Details struct {
	BkLoss  string `json:"bkLoss"`
	BkPx    string `json:"bkPx"`
	Ccy     string `json:"ccy"`
	PosSide string `json:"posSide"`
	Side    string `json:"side"`
	Sz      string `json:"sz"`
	Ts      string `json:"ts"`
}
type Data struct {
	Details    []Details `json:"details"`
	InstFamily string    `json:"instFamily"`
	InstID     string    `json:"instId"`
	InstType   string    `json:"instType"`
	Uly        string    `json:"uly"`
}
