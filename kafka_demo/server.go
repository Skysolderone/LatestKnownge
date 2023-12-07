package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	//配置
	config := sarama.NewConfig()
	//设置属性
	config.Producer.Timeout = 5 * time.Second
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//construct msg
	// msg := sarama.ProducerMessage{}
	// msg.Topic = "web_log"
	// msg.Value = sarama.StringEncoder("this is a test log")
	//connect kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	//send msg
	// pid, offset, err := client.SendMessage(&msg)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%v,%v", pid, offset)
	for i := 0; i < 10; i++ {
		//create message
		value := fmt.Sprintf("sync message ,index=%d", i)
		msg := &sarama.ProducerMessage{
			Topic: "topic1",
			Value: sarama.ByteEncoder(value),
		}
		part, offset, err := client.SendMessage(msg)
		if err != nil {
			log.Printf("send message error:%s \n", err.Error())
		} else {
			fmt.Printf("SUCCESS: value=%s, partition=%d, offset=%d \n", value, part, offset)
		}
		time.Sleep(time.Second * 2)
	}
}
