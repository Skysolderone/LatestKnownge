package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//construct msg
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	//connect kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	//send msg
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v,%v", pid, offset)
}
