package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

const ServerName = "HelloService"

func DialHelloService(network string, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatal(err)
	}
	// client := rpc.NewClient(conn)
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{Client: client}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(ServerName+".Hello", request, reply)
}

// 编译期间检查
// var _ HelloServiceInterface = (*HelloServiceClient)(nil)
type clientRequest struct {
	Method string `json:"method"`
	Params [1]any `json:"params"`
	Id     uint64 `json:"id"`
}

func main() {
	// conn, err := net.Dial("tcp", ":4567")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client := rpc.NewClient(conn)
	client, err := DialHelloService("tcp", ":4567")
	if err != nil {
		log.Fatal(err)
	}
	var reply string
	// err = client.Call("HelloService.Hello", "hello", &reply)
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
