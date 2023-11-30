package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signalchannel := make(chan os.Signal, 1)
	signal.Notify(signalchannel, syscall.SIGINT) //可添加多个信号 除了无法忽略的信号  kill SG
	go func() {
		<-signalchannel
		fmt.Println("close")
		os.Exit(0)
	}()
}
