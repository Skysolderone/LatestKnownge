package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	_, err = conn.Write([]byte("heelo wws"))
	if err != nil {
		log.Fatal(err)
	}
}
