package svc

import (
	"v1/api/internal/config"
	"v1/rpc/add/adder"
	"v1/rpc/check/checker"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder
	Checker checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
