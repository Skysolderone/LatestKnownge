package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
)

func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}

func main() {
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}
	wsurl := "wss://fstream.binance.com/ws/!forceOrder@arr"
	conn, _, err := Dialer.Dial(wsurl, nil)
	if err != nil {
		fmt.Println(err)
	}
	conn.SetReadLimit(655350)
	go func() {
		keepAlive(conn, time.Second*60)
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
			}
			data := ForceOrderEvent{}
			sonic.Unmarshal(msg, &data)
			fmt.Printf("%#v", data)

		}
	}()
	for {
	}
}

type ForceOrderEvent struct {
	EventType string      `json:"e"`
	EventTime int64       `json:"E"`
	Order     OrderDetail `json:"o"`
}

// 定义订单详细信息的结构体
type OrderDetail struct {
	Symbol             string `json:"s"`
	Side               string `json:"S"`
	OrderType          string `json:"o"`
	TimeInForce        string `json:"f"`
	Quantity           string `json:"q"`
	Price              string `json:"p"`
	AveragePrice       string `json:"ap"`
	OrderStatus        string `json:"X"`
	LastFilledQuantity string `json:"l"`
	CumulativeQuantity string `json:"z"`
	TransactionTime    int64  `json:"T"`
}
