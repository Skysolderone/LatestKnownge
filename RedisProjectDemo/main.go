package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

// 秒杀处理
func handledSeckKill(userId string) error {
	const (
		productId = "product_001"
		lockKey   = "seckkill_lock:" + productId

		lockDuration = 3 * time.Second
	)
	lockValue := userId
	// 分布式锁
	ok, err := redisClient.SetNX(context.Background(), lockKey, lockValue, lockDuration).Result()
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("%s mutex failed", userId)
	}
	defer redisClient.Del(context.Background(), lockKey)
	// 模拟秒杀逻辑处理时间
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	// check 库存
	inventory, err := redisClient.Get(context.Background(), "inventory:"+productId).Int()
	if err != nil {
		return err
	}
	if inventory == 0 {
		return fmt.Errorf("库存不足")
	}

	// 扣减库存
	err = redisClient.Decr(context.Background(), "inventory:"+productId).Err()
	if err != nil {
		return err
	}
	fmt.Printf("%s success", userId)
	return nil
}

// 主函数并发处理秒杀请求
func main() {
	const (
		numUsers   = 1000
		numSeconds = 5
	)
	var wg sync.WaitGroup
	wg.Add(numUsers)
	for i := 0; i < numUsers; i++ {
		go func(userid int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			// 处理秒杀逻辑
			err := handledSeckKill(fmt.Sprintf("user_%d", userid))
			if err != nil {
				log.Fatal(err)
			}
		}(i)
	}
}
