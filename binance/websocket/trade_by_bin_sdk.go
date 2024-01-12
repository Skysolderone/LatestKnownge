package main

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
)

func main() {

	r := iris.New()
	r.Any("/testBinWs", wsTest)
	r.Listen(":8080")
}

var upg = websocket.Upgrader{}

func wsTest(ctx iris.Context) {
	c, err := upg.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		obj, err := sonic.Marshal(event)
		if err != nil {
			fmt.Println(err)
		}
		c.WriteMessage(1, obj)

	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsKlineServe("icpusdt", "1m", wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
