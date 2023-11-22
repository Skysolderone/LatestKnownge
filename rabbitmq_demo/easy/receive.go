package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:15672/")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := ch.Consume(
		q.Name, "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	var forver chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("receive :%s", d.Body)
		}
	}()
	log.Println("finish")
	<-forver
}
