package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()
	//处理订单队列
	for {
		order, err := client.RPop(context.Background(), "order_queue").Result()
		if err == redis.Nil {
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			log.Println(err)
			continue
		}
		processOrder(order)

	}
}
func processOrder(order string) {
	fmt.Printf("Processing order:%s\n", order)
	//process
}
