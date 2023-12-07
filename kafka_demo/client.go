package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/IBM/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

// basic
func main1s() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		panic(err)
	}
	partitionList, err := consumer.Partitions("web_logs")
	if err != nil {
		panic(err)
	}
	fmt.Println(partitionList)
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("%d,%d,%v,%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
}

// level1
func main() {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = -2
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Group.Return.Notifications = true

	brokers := []string{"127.0.0.1:9092"}
	topics := []string{"topic1"}
	consumer, err := cluster.NewConsumer(brokers, "consumer-group", topics, config)
	if err != nil {
		fmt.Printf("new consumer error :%s\n", err.Error())
		return
	}
	defer consumer.Close()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		for err := range consumer.Errors() {
			fmt.Printf("consumer error:%s \n", err.Error())
		}
	}()
	go func() {
		for ntf := range consumer.Notifications() {
			fmt.Printf("consumer notifications error :%v \n", ntf)
		}
	}()

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Printf("%s/%d/%dt%st%sn", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // 上报offset
			} else {
				fmt.Println("监听服务失败")
			}
		case <-signals:
			return
		}
	}
}
