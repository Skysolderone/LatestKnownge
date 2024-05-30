package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// IntervalSecond := map[string]uint{"1m": 60, "5m": 300, "15m": 900, "30m": 1800, "1h": 3600, "2h": 7200, "4h": 14400, "8h": 28800, "1d": 86400, "4H": 14400, "1H": 3600, "2H": 7200, "8H": 28800, "1D": 86400}
	// str := "15m"
	// if _, ok := IntervalSecond[str]; ok {
	// 	fmt.Println("SUCCEESS")
	// }
	num := 1717037099.999
	str := strconv.FormatFloat(num, 'f', -1, 64)
	fmt.Println(str)
	change := strings.Split(str, ".")
	fmt.Println(change[0])
}
