package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// 需要本地docker redis
func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //no password set
		DB:       0,  //use default db
	})
	//basic
	ctx := context.Background()
	// pong, err := client.Ping(ctx).Result()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(pong)
	//key value
	client.Set(ctx, "key", "i am value", time.Hour)
	result := client.Get(ctx, "key")
	log.Println(result)
	defer client.Close()
	//生成订单队列
	for i := 1; i <= 10; i++ {
		order := fmt.Sprintf("Order %d", i)
		_, err := client.LPush(context.Background(), "order_queue", order).Result()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("generate %s\n", order)
		time.Sleep(1 * time.Second)
	}
}
