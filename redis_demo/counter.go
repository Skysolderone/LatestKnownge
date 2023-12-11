package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	// client.Set(ctx, "counter", 1, 0)
	err := client.Incr(ctx, "counter").Err()
	if err != nil {
		log.Fatal(err)
	}
	value := client.Get(ctx, "counter")
	fmt.Println(value)
}
