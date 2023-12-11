package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "my-topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	handleErr(err)
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(1, 1e6) //fetch min 1b  max 1mb
	b := make([]byte, 10e3)         //10kb
	for {
		n, err := batch.Read(b)
		if err != nil {
			//log.Println(err)
			break
		}
		//fmt.Println(b)
		fmt.Println(string(b[:n]))
	}
	err = batch.Close()
	handleErr(err)
	err = conn.Close()
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
