package main

import (
	"fmt"
	"time"
)

func main() {
	nows := time.Now().Unix()
	// fmt.Println(now)
	now := time.Unix(nows, 0)
	end_time := time.Unix(1721205245, 0)
	times := end_time.Sub(now)
	fmt.Println(int(times.Hours() / 24))
}
