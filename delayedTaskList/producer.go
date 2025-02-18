package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func AddDelayedTask(ctx context.Context, taskID string, delay time.Duration) error {
	t := time.Now().Add(delay).Unix()
	return rdb.ZAdd(ctx, "delayed_tasks", redis.Z{
		Score:  float64(t),
		Member: taskID,
	}).Err()
}

func main() {
	ctx := context.Background()
	// pong, err := rdb.Ping(context.Background()).Result()
	err := AddDelayedTask(ctx, "task2", 10*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println("producer success")
}
