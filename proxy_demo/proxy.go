package main

import (
	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()
	// target, err := net.Dial("tcp", "8.218.201.224:80")
	target, err := net.Dial("tcp", "192.168.10.128:80")
	if err != nil {
		log.Println("connect:", err)
		return
	}
	defer target.Close()
	go func() {
		_, err := io.Copy(target, conn)
		if err != nil {
			log.Println("copy:", err)
		}
	}()
	_, err = io.Copy(conn, target)
	if err != nil {
		log.Println("copy2:", err)
	}
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:55")
	if err != nil {
		log.Fatal("listen:", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("accept:", err)
			continue
		}
		go handle(conn)
	}
}
