package main

import (
	"log"
	"net"

	"v1/internal/client"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client.Start(conn)
}
