package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WeightRandom struct {
	servers []string
	weights []int
}

func NewWeigthRandom(servers []string, weights []int) *WeightRandom {
	return &WeightRandom{
		servers: servers,
		weights: weights,
	}
}
func (wr *WeightRandom) GetWeightRandom() string {
	rand.Seed(time.Now().UnixNano())
	toTotalWeight := 0
	for _, weight := range wr.weights {
		toTotalWeight += weight
	}
	randWight := rand.Intn(toTotalWeight)
	for i, weight := range wr.weights {
		if randWight < weight {
			return wr.servers[i]
		}
		randWight -= weight
	}
	return wr.servers[len(wr.servers)-1]
}

// 加权随机法
func main() {
	servers := []string{"String1", "String2", "String3"}
	weights := []int{2, 1, 3}
	wr := NewWeigthRandom(servers, weights)
	for i := 1; i < 10; i++ {
		fmt.Println("Request sent to", wr.GetWeightRandom())
	}
}
