package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

// async
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_10_0_1
	fmt.Println("start make producer")
	producer, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		log.Printf("new async producer error:%s \n", err.Error())
		return
	}
	defer producer.AsyncClose()
	fmt.Println("start goroutine")
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case suc := <-p.Successes():
				fmt.Println("offset:", suc.Offset, "timesamp:", suc.Timestamp.String(), "partitions:", suc.Partition)
			case fail := <-p.Errors():
				fmt.Println("error:", fail.Error())
			}
		}

	}(producer)
	var value string
	for i := 0; ; i++ {
		time.Sleep(2 * time.Second)
		value = fmt.Sprintf("async message,index=%d", i)
		msg := &sarama.ProducerMessage{
			Topic: "topic1",
			Value: sarama.ByteEncoder(value),
		}
		producer.Input() <- msg
	}
}
