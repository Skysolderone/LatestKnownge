package main

import (
	"log"
	"time"
)

func main() {
	timer := time.NewTicker(time.Second * 5)
	// timer2 := time.NewTicker(time.Second * 8)
	// timer3 := time.NewTicker(time.Second * 1)
	// timer3.Stop()
	log.Println(5)
	for {
		select {
		case <-timer.C:
			log.Println(5)
			timer.Reset(time.Second * 1)
			// timer.Stop()
			// case <-timer2.C:
			// 	log.Println(5)
			// 	timer2.Stop()
			// 	timer3.Reset(time.Second * 1)
			// case <-timer3.C:
			// 	log.Println(1)
		}
	}
}
