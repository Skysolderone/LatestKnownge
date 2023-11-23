package main

import (
	"context"
	"log"
	"net"
	"v1/demo"

	grpc "google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *demo.Request) (*demo.Response, error) {
	return &demo.Response{Message: "Hello," + req.Message}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	demo.ResisterDemoServiceServer(srv, &server{})
	if err := srv.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
