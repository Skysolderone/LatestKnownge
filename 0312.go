package main

import (
	"log"
	"time"
)

func main() {
	diff := time.Now().Sub(time.Now().Truncate(time.Hour))
	nice := time.Since(time.Now().Truncate(time.Hour))
	log.Println(diff)
	log.Println(nice)

	timer := time.NewTicker(time.Second * 3)
	for range timer.C {
		log.Println(3)
		timer.Reset(time.Second * 5)

	}
}
