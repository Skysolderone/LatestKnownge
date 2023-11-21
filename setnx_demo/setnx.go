package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// 基于redis实现的分布式锁  用于分布式系统
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
func incr() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	var lockkey = "counter_lock"
	var counterKey = "counter"
	//lock
	resp := client.SeNx(ctx, lockkey, 1, time.Second*5)
	lockSuccess, err := resp.Result()
	if err != nil || !lockSuccess {
		fmt.Println(err, " locak result ", lockSuccess)
		return
	}
	//counter++
	getResp := client.Get(ctx, counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := client.Set(ctx, counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			log.Println("set value error:", err)
		}

	}
	fmt.Println("cnt value :", cntValue)
	delResp := client.Del(ctx, lockkey)
	unlockSuccess, err := delResp.Result()
	if err == nil || unlockSuccess > 0 {
		log.Println("unlock success")
	} else {
		log.Println("unlock failed")
	}
}
