package main

import (
	"context"
	"fmt"
	"log"
	"v1/demo"

	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("loacalhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := demo.NewDemoServiceClient(conn)
	req := &demo.Request{Message: "Gopher"}
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response from server %s\n", resp.Message)
}
