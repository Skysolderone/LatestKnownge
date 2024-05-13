package svc

import (
	"v1/api/internal/config"
	"v1/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Transformer transformer.Transformer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),
	}
}
