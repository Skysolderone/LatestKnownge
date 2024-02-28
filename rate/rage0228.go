package main

import (
	"fmt"

	"golang.org/x/time/rate"
)

func main() {
	s := rate.Sometimes{Every: 2}
	s.Do(func() { fmt.Println("1") })
	s.Do(func() { fmt.Println("2") })
	s.Do(func() { fmt.Println("3") })
}
