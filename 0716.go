package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	s := "1721120400000"
	str, _ := strconv.ParseInt(s, 10, 64)
	tim := time.Unix(str/1000, 0)
	fmt.Println(tim.Format("2006-01-02 15:04:05"))
}
