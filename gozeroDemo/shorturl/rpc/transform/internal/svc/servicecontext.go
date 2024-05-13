package svc

import (
	"v1/rpc/transform/internal/config"
	"v1/rpc/transform/transformer/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c     config.Config
	Model model.ShorturlModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:     c,
		Model: model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
