package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BucketSort(s []int, size int) {
	minValue := s[0]
	maxValue := s[0]
	for i := 0; i < len(s); i++ {
		if minValue > s[i] {
			minValue = s[i]
		}
		if maxValue < s[i] {
			maxValue = s[i]
		}
	}
	fmt.Println(minValue)
	fmt.Println(maxValue)
	bucket := make([][]int, (maxValue-minValue)/size+1)
	for i := 0; i < len(s); i++ {
		bucket[(s[i]-minValue)/size] = append(bucket[(s[i]-minValue)/size], s[i])
	}

	key := 0
	for _, bucketnum := range bucket {
		if len(bucketnum) <= 0 {
			continue
		}
		fmt.Println(bucket)
		insertSort(bucketnum)
		for _, value := range bucketnum {
			s[key] = value
			key = key + 1
		}

	}
	return
}
func insertSort(s []int) {

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	num := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(num); i++ {
		num[i] = rand.Intn(100)
	}
	fmt.Println(num)
	BucketSort(num, 10)
	fmt.Println(num)
}
