package main

import (
	"log"
	"strconv"

	"github.com/nsqio/go-nsq"
)

func main() {
	var topic = "learnnsq"
	config := nsq.NewConfig()
	w, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal("error:", err)
	}
	for i := 0; i < 100; i++ {
		err := w.Publish(topic, []byte("nsq wiht golang step:"+strconv.Itoa(i)))
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Stop()
}
