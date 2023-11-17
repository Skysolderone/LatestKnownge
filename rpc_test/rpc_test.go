package test

import (
	"net"
	"net/rpc"
	"testing"
)

// basic
// type Arith struct{}
// type Args struct {
// 	A, B int
// }

func (t *Arith) Sum(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func TestRpcServer(t *testing.T) {
	arith := new(Arith)
	rpc.Register(arith)
	liseter, err := net.Listen("tcp", ":1234")
	if err != nil {
		t.Log(err)
	}
	t.Log("Rpc Server listen 1234")
	for {
		conn, err := liseter.Accept()
		if err != nil {
			t.Log(err)
			continue
		}
		go rpc.ServeConn(conn)

	}
}

func TestClient(t *testing.T) {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		t.Log(err)
	}
	defer client.Close()
	args := Args{A: 7, B: 8}
	var reply int
	err = client.Call("Arith.Sum", args, &reply)
	if err != nil {
		t.Log(err)
	}
	t.Log(args.A, args.B, reply)
}
