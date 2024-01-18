package main

import (
	"fmt"
	"sync"
)

type WeightedRoundRobin struct {
	servers    []string
	weights    []int
	currentInx int
	lock       sync.Mutex
}

func NewWeightedRoundRobin(servers []string, weights []int) *WeightedRoundRobin {
	return &WeightedRoundRobin{
		servers:    servers,
		weights:    weights,
		currentInx: 0,
	}

}
func (wrr *WeightedRoundRobin) GetNextServer() string {
	wrr.lock.Lock()
	defer wrr.lock.Unlock()
	server := wrr.servers[wrr.currentInx]
	wrr.currentInx = (wrr.currentInx + 1) % len(wrr.servers)
	return server
}
//加权轮询
func main() {
	servers := []string{"Server1", "Server2", "Server3"}
	weights := []int{1, 2, 3}
	wrr := NewWeightedRoundRobin(servers, weights)
	for i := 0; i < 10; i++ {
		fmt.Println("Request send to:", wrr.GetNextServer())
	}
}
