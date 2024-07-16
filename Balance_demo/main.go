package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"v1/balance"
)

func main() {
	var insts []*balance.Instance
	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port, _ := strconv.Atoi(fmt.Sprintf("880%d", i))
		one := balance.NewInstance(host, port)
		insts = append(insts, one)
	}
	name := "random"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	for {
		inst, err := balance.DoBalance(name, insts)
		if err != nil {
			fmt.Println("do balance err")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
