package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

type myHandler struct{}

func (*myHandler) HandleMessage(message *nsq.Message) error {
	fmt.Println(string(message.Body))
	return nil
}
func main() {
	var topic = "learnnsq"
	var wg sync.WaitGroup
	wg.Add(1)
	config := nsq.NewConfig()
	q, err := nsq.NewConsumer(topic, "channel", config)
	if err != nil {
		log.Fatal(err)
	}
	q.AddHandler(&myHandler{})
	err1 := q.ConnectToNSQD("127.0.0.1:4150")
	if err1 != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
