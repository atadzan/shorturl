package svc

import (
	"github.com/atadzan/shorturl/internal/config"
	"github.com/atadzan/shorturl/rpc/transform/transformer"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer // manual code
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)), // manual code
	}
}
