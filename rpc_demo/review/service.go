package main

import (
	"encoding/json"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const serverName = "HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServiceInterface) error {
	return rpc.RegisterName(serverName, srv)
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	log.Println("HelloService Hello")
	*reply = "service receive:" + request
	return nil
}

type serverRequest struct {
	Method string           `json:"method"`
	Params *json.RawMessage `json:"params"`
	Id     *json.RawMessage `json:"id"`
}

// 注册服务
func main() {
	// err := rpc.RegisterName("HelloService", new(HelloService))

	// if err != nil {
	// 	log.Fatal(err)
	// }
	_ = RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":4567")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
