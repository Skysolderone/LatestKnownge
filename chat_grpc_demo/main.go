package main

import (
	"log"
	"net"
	proto "v1/gen"
	"v1/handler"

	"google.golang.org/grpc"
)

func main() {
	grpcserver := grpc.NewServer()
	var conn []*handler.Connection
	pool := &handler.Pool{
		Connection: conn,
	}
	proto.RegisterBroadcastServer(grpcserver, pool)
	lisetener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := grpcserver.Serve(lisetener); err != nil {
		log.Fatal(err)
	}
}
