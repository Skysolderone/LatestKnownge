package main

import (
	"fmt"
	"log"
	"time"

	"github.com/maypok86/otter"
)

func main() {
	cache, err := otter.MustBuilder[string, string](10_000).
		CollectStats().
		Cost(func(key, value string) uint32 {
			return 1
		}).
		WithTTL(time.Hour).Build()
	if err != nil {
		log.Fatal(err)
	}
	cache.Set("key", "value0406")
	value, Ok := cache.Get("key")
	if !Ok {
		log.Fatal("NOT EXISTS")
	}
	fmt.Println(value)
	cache.Close()
	//   // 设置不同TTL的值
	// 上面使用  WithVariableTTL().
	// cache.Set("key1", "value1", time.Hour)   // 1小时
	// cache.Set("key2", "value2", time.Minute) // 1分钟

	// value, ok := cache.Get("key1") // 从缓存中获取值
	// if !ok {
	//     panic("未找到对应的键")
	// }
	// fmt.Println(value) // 输出获取到的值
	// cache.Close() // 关闭缓存实例
}
