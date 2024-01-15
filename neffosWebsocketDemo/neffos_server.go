package main

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

var wsServer = websocket.Namespaces{
	"default": websocket.Events{
		websocket.OnNamespaceConnected: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			ctx := websocket.GetContext(nsConn.Conn)
			log.Printf("[%s] connect to namespace [%s] with IP [%s]", nsConn, msg.Namespace, ctx.RemoteAddr())
			return nil

		},
		websocket.OnNamespaceDisconnect: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("[%s] disconnect namspace [%s]", &nsConn, msg.Namespace)
			return nil
		},
		"chat":func(nsConn *websocket.NSConn,msg websocket.Message)error{
			log.Printf("%s send %s",nsConn,string(msg.Body))
			nsConn.Conn.Server().Broadcast(nsConn,msg)
			return nil
		},
	},
}

func main() {
	app := iris.New()
	websocketServer := websocket.New(
		websocket.DefaultGorillaUpgrader,
		wsServer)
	app.Get("/", websocket.Handler(websocketServer))
	app.Listen(":8080")
}
