package main

import (
	"context"
	"fmt"

	"github.com/bitfinexcom/bitfinex-api-go/pkg/models/ticker"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
)

func main() {
	clt := websocket.New()
	// clt.SubscribeTicker(context.Background(), "XAUT:UST")
	clt.SubscribeTicker(context.Background(), "tXAUT:UST")
	clt.Connect()

	for {
		select {
		case crash := <-clt.Listen():
			switch v := crash.(type) {
			case *ticker.Ticker:
				// log.Printf("%T: %+v\n", v, v)
				symbolname := v.Symbol
				fmt.Printf("---------------%s: %+v\n", symbolname, v)
			}
		}
	}
}
