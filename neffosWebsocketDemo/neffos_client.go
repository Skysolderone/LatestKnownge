package main

import (
	"context"
	"log"
	"time"

	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

var clientEvents = websocket.Namespaces{
	"default": websocket.Events{
		websocket.OnNamespaceConnect: func(c *websocket.NSConn, msg websocket.Message) error {
			log.Printf("connect Namespace [%s]", msg.Namespace)
			return nil
		},
		websocket.OnNamespaceDisconnect: func(c *websocket.NSConn, msg websocket.Message) error {
			log.Printf("disconnect Namespace [%s]", msg.Namespace)
			return nil
		},
		"chat": func(c *websocket.NSConn, msg websocket.Message) error {
			log.Printf("%s", string(msg.Body))
			return nil
		},
	},
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()
	dialer := websocket.DefaultGobwasDialer
	//client, err := websocket.Dial(ctx, dialer, "ws://localhost:8080/", clientEvents)
	client, err := websocket.Dial(ctx, dialer, "ws://8.222.221.115:14000/binanceus", clientEvents)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	client.Connect(ctx, "default")
	if err != nil {
		panic(err)
	}

	// fmt.Fprint(os.Stdout, ">> ")
	// scanner := bufio.NewScanner(os.Stdin)
	for {
		clientEvents.On("default", "BTCUSDT", func(conn *neffos.NSConn, message neffos.Message) error {
			log.Println(string(message.Body))
			return nil
		})
		// if !scanner.Scan() {
		// 	log.Printf("ERROR: %v", scanner.Err())
		// 	return
		// }
		// text := scanner.Bytes()
		// if bytes.Equal(text, []byte("exit")) {
		// 	if err := c.Disconnect(nil); err != nil {
		// 		log.Printf("reply from server :%v", err)
		// 	}
		// 	break
		// }
		// if ok := c.Emit("chat", text); !ok {
		// 	break
		// }
		// fmt.Fprint(os.Stdout, ">> 	")
	}

}
