package main

import (
	"log"
	"net"

	"v1/internal/server"
	"v1/pkg/pow"
	"v1/pkg/wisdom"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	srv := server.NewServer(pow.PoWImpl{}, wisdom.WisdomImpl{})
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go srv.Handle(conn)
	}
}
