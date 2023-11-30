package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// telnet 协议
func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	// defer listener.Close()
	//优雅关闭
	signalchannel := make(chan os.Signal, 1)
	signal.Notify(signalchannel, syscall.SIGINT)
	go func() {
		<-signalchannel
		fmt.Println("close")
		listener.Close()
		os.Exit(0)
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handler(conn)
	}
}
func handler(conn net.Conn) {
	defer conn.Close()
	// buffer := make([]byte, 1024)
	// for {
	// 	n, err := conn.Read(buffer)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	data := buffer[:n]
	// 	fmt.Println("receive :", string(data))
	// 	_, err = conn.Write(data)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	//telnet 回音服务器
	fmt.Fprint(conn, "welcome\n")
	scaner := bufio.NewScanner(conn)
	for scaner.Scan() {
		input := scaner.Text()
		fmt.Println("receve:", input)
		fmt.Fprintln(conn, "回音:"+input)
	}
	if err := scaner.Err(); err != nil {
		log.Fatal(err)
	}
}
