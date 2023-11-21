package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-zookeeper/zk"
)

// 基于 ZooKeeper 的锁与基于 Redis 的锁的不同之处在于 Lock 成功之前会一直阻塞，
// 这与我们单机场景中的 mutex.Lock 很相似。
func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
	if err != nil {
		panic(err)
	}
	l := zk.NewLock(c, "/lock", zk.WorldCL(zk.PermAll))
	err = l.lock()
	if err != nil {
		panic(err)
	}
	fmt.Println("lock succ")
	time.Sleep(time.Second * 10)
	//do something
	log.Println("hello zk ")
	l.unlock()
	fmt.Println("unlock succ")
}
