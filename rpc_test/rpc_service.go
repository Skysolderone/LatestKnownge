package main

import (
	"net"
	"net/rpc"
	"v1/rpcmodel"
)

func main() {
	arith := new(rpcmodel.Arith)
	rpc.Register(arith)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		rpc.ServeConn(conn)
	}

}
