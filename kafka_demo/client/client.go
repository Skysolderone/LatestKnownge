package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "my-topic", 0)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
