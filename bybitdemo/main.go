package main

import (
	"fmt"

	"github.com/bytedance/sonic"

	gw "github.com/gorilla/websocket"
)

//	{
//	    "req_id": "test", // 可選
//	    "op": "subscribe",
//	    "args": [
//	        "orderbook.1.BTCUSDT"
//	    ]
//	}
type req struct {
	Req_id string   `json:"req_id"`
	Op     string   `json:"op"`
	Args   []string `json:"args"`
}

func main() {
	// url := "wss://stream.bybit.com/v5/public/linear"
	// c, _, err := websocket.DefaultDialer.Dial(url, nil)
	// if err != nil {
	// 	panic(err)
	// }
	// defer c.Close()
	// s := req{
	// 	Req_id: "12312312",
	// 	Op:     "subscribe",
	// 	Args:   []string{"tickers.GPSUSDT"},
	// }
	// data, _ := sonic.Marshal(s)
	// err = c.WriteMessage(websocket.TextMessage, data)
	// if err != nil {
	// 	panic(err)
	// }
	// time.AfterFunc(time.Second*5, func() {
	// 	s := req{
	// 		Req_id: "3b88e279-fedd-4d80-b1ac-d0bfcc83e066",
	// 		Op:     "subscribe",
	// 		Args:   []string{"tickers.SHELLUSDT"},
	// 	}
	// 	data, _ := sonic.Marshal(s)
	// 	err = c.WriteMessage(websocket.TextMessage, data)
	// })
	// for {
	// 	_, data, err := c.ReadMessage()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(string(data))
	// }
	checkBitgetSymbol(true, "BSSUSDT")
}

func checkBitgetSymbol(futures bool, symbol string) bool {
	url := "wss://ws.bitget.com/v2/ws/public"
	typ := "futures"
	arg := Args{InstType: "USDT-FUTURES", Channel: "ticker", InstId: symbol}
	s := Sub{
		Op:   "subscribe",
		Args: []Args{arg},
	}
	if !futures {
		typ = "spot"
		url = "wss://ws.bitget.com/v2/ws/public"
		arg = Args{InstType: "SPOT", Channel: "ticker", InstId: symbol}
		s = Sub{
			Op:   "subscribe",
			Args: []Args{arg},
		}
	}

	c, _, err := gw.DefaultDialer.Dial(url, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	data, _ := sonic.Marshal(s)
	err = c.WriteMessage(gw.TextMessage, data)
	if err != nil {
		// log.Error("checksymbol", err)
		fmt.Println(err)
	}

	for {
		_, data, err := c.ReadMessage()
		if err != nil {
			// log.Error("checksymbol read data", err)
			fmt.Println(err)
		}
		fmt.Println(string(data), typ)

	}
}

type Sub struct {
	Op   string `json:"op"`
	Args []Args `json:"args"`
}
type Args struct {
	InstType string `json:"instType"`
	Channel  string `json:"channel"`
	InstId   string `json:"instId"`
}
