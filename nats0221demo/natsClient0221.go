package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	sub, err := js.PullSubscribe("wws-project", "wws-consumer", nats.PullMaxWaiting(1024))
	if err != nil {
		log.Fatal(err)
	}
	msg, err := sub.NextMsg(time.Second * 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg.Data)
}
