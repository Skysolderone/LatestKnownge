package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Random struct {
	servers []string
}

func NewRandom(servers []string) *Random {
	return &Random{
		servers: servers,
	}
}
func (r *Random) GetRandomServer() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(r.servers))
	return r.servers[index]
}
//随机
func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	random := NewRandom(servers)
	for i := 0; i < 10; i++ {
		fmt.Println("Request send to:", random.GetRandomServer())
	}
}
