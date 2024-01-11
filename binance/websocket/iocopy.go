package main

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upg = websocket.Upgrader{}

//  尝试使用io.copy 直接传输两个ws的数据
//  需要使用ws.UnderlyingConn() 来进行数据拷贝
func main() {
	r := gin.Default()
	r.GET("/tesws", testWs)
	r.Run()
}
func testWs(c *gin.Context) {
	api_url := "wss://stream.binance.us:9443/ws/icpusdt@kline_1m"
	cs, _, err := websocket.DefaultDialer.Dial(api_url, nil)
	if err != nil {
		panic(err)
	}
	ws, err := upg.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		defer ws.Close()
		defer cs.Close()
		for {
			_, err := io.Copy(ws.UnderlyingConn(), cs.UnderlyingConn())
			if err != nil {
				return
			}
		}
	}()
}
