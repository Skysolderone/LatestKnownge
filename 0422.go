package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	c := strconv.Itoa(rand.Intn(1<<31 - 1))
	fmt.Println(c)
}
