package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 本地没有rabserver 需使用docker
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:15672/")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Println(err)
	}
	q, err := channel.QueueDeclare(
		"hello", false, false, false, false, nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	body := "hello wws"
	err = channel.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("send body:", body)
}
