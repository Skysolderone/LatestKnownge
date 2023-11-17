package main

import (
	"fmt"
	"net/rpc"
	"v1/rpcmodel"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	args := &rpcmodel.Args{7, 8}
	var reply int
	err = client.Call("Arith.Sum", args, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(args.A, args.B, reply)

}
