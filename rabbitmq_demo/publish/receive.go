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
	err = ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	// err = ch.Qos(1, 0, false)
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(q.Name, "", "logs", false, nil)
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
			// dotCount := bytes.Count(d.Body, []byte("."))
			// t := time.Duration(dotCount)
			// time.Sleep(t * time.Second)
			// log.Println("DONE")
			// d.Ack(false)
		}
	}()
	log.Println("waiting msg")
	<-forver
}
