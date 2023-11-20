package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 本地没有rabserver 需使用docker
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	channel, err := conn.channel()
	if err != nil {
		log.Println(err)
	}

	

}
