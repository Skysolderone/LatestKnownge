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
}
