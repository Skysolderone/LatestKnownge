package main

import (
	"fmt"
	"time"

	"github.com/sqids/sqids-go"
)

// 构建sql unique id
func main() {
	s, _ := sqids.New()
	id, _ := s.Encode([]uint64{1234567890})
	fmt.Println(id)
	start := time.Now().Unix()
	end := time.Now().Add(24 * time.Hour).Unix()

	id, _ = s.Encode([]uint64{uint64(start), uint64(end)}) // "s6eUn008oGU27p"
	fmt.Println(id)

	numbers := s.Decode(id) // [1714879533 1714965933]
	fmt.Println(numbers)
}
