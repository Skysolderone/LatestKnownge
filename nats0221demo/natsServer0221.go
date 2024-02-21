package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

// go get github.com/nats-io/nats.go
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
	msg := []byte("Hello,NATS jetstream!")
	_, err = js.Publish("wws-project", msg)
	if err != nil {
		log.Fatal(err)
	}
		
}
