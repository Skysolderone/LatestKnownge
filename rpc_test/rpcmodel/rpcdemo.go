package rpcmodel

type Arith struct{}
type Args struct {
	A, B int
}

func (a *Arith) Sum(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
