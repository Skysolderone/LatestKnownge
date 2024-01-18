package main

import (
	"fmt"
	"sync"
)

type Roundrobin struct {
	servers []string
	index   int
	lock    sync.Mutex
}

func NewRoundrobin(server []string) *Roundrobin {
	return &Roundrobin{
		servers: server,
		index:   0,
	}
}
func (rr *Roundrobin) GetNextServer() string {
	rr.lock.Lock()
	defer rr.lock.Lock()
	server := rr.servers[rr.index]
	rr.index = (rr.index + 1) % len(rr.servers)
	return server
}
//轮询
func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	rr := NewRoundrobin(servers)
	for i := 0; i < 10; i++ {
		fmt.Println("Request send to :", rr.GetNextServer())
	}
}
