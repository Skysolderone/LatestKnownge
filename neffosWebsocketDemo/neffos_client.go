package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kataras/iris/v12/websocket"
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
			log.Printf("%s send %s", c, string(msg.Body))
			return nil
		},
	},
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()
	dialer := websocket.DefaultGobwasDialer
	client, err := websocket.Dial(ctx, dialer, "ws://localhost:8080/", clientEvents)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	c, err := client.Connect(ctx, "default")
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, ">> ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			log.Printf("ERROR: %v", scanner.Err())
			return
		}
		text := scanner.Bytes()
		if bytes.Equal(text, []byte("exit")) {
			if err := c.Disconnect(nil); err != nil {
				log.Printf("reply from server :%v", err)
			}
			break
		}
		if ok := c.Emit("chat", text); !ok {
			break
		}
		fmt.Fprint(os.Stdout, ">> 	")
	}

}
