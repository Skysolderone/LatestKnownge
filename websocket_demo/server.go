package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upg = websocket.Upgrader{}

func main() {
	r := gin.Default()
	r.GET("/ws", func(ctx *gin.Context) {
		conn, err := upg.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		log.Println("CONNECT")
	})
	r.Run()
}
