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

func StartConsumer(ctx context.Context) {
	for {
		t := time.Now().Unix()
		tasks, err := rdb.ZRangeByScore(ctx, "delayed_tasks", &redis.ZRangeBy{
			Min:    "0",
			Max:    fmt.Sprintf("%d", t),
			Offset: 0,
			Count:  10,
		}).Result()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// fmt.Println(tasks)
		if len(tasks) == 0 {
			// fmt.Println("zero")
			continue
		}
		// 使用 Lua 脚本原子化移除任务并返回任务列表
		luaScript := `
			local tasks = redis.call('ZRANGEBYSCORE', KEYS[1], '-inf', ARGV[1])
			if #tasks > 0 then
				redis.call('ZREM', KEYS[1], unpack(tasks))
			end
			return tasks
		`
		keys := []string{"delayed_tasks"}
		vals := []interface{}{t}

		res, err := rdb.Eval(ctx, luaScript, keys, vals).Result()
		if err != nil {
			fmt.Println("执行Lua脚本失败:", err)
			continue
		}

		// 处理任务
		if res != nil {
			taskIDs := res.([]interface{})
			for _, taskID := range taskIDs {
				fmt.Printf("执行任务: %s\n", taskID)
				// TODO: 执行具体任务逻辑
			}
		}
	}
}

func main() {
	fmt.Println("run consumer")
	ctx := context.Background()
	go StartConsumer(ctx)
	select {}
}
