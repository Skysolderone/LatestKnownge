package main

import (
	"fmt"
	"strings"
)

var m = make(map[string]string)

func main() {
	sort := "vip desc"
	switch {
	case strings.Contains(sort, "total_money") && strings.Contains(sort, "asc"):
		fmt.Println("1")
	case strings.Contains(sort, "total_money") && strings.Contains(sort, "desc"):
		fmt.Println("2")
	case strings.Contains(sort, "vip") && strings.Contains(sort, "asc"):
		fmt.Println("3")
	case strings.Contains(sort, "vip") && strings.Contains(sort, "desc"):
		fmt.Println("4")
	case strings.Contains(sort, "money") && strings.Contains(sort, "asc"):
		fmt.Println("5")
	case strings.Contains(sort, "money") && strings.Contains(sort, "desc"):
		fmt.Println("6")

	}
}
