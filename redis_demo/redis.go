package main

import (
	"context"
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
}
