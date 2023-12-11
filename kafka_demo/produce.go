package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "my-topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	handleErr(err)
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second)) //写的超时时间
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("One")},
		kafka.Message{Value: []byte("Two")},
		kafka.Message{Value: []byte("Three")})
	handleErr(err)
	err = conn.Close()
	handleErr(err)

}
func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
