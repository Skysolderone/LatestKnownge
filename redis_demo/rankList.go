package main

import (
	"fmt"
	"sort"
	"time"
)

var scores []int64

// basic
func main() {
	sorts()
}

func sorts() {
	now := time.Now().UnixMilli()
	for i := 0; i < 10; i++ {
		scores = append(scores, toScores(int64(i+10000), now))
	}
	fmt.Println(scores)
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})
	for _, score := range scores {
		fmt.Println(load(score))
	}
}

func toScores(i int64, time int64) int64 {
	var score int64
	score = (score | i) << 41
	score = score | time
	return score

}
func load(i int64) int64 {
	return i >> 41
}
