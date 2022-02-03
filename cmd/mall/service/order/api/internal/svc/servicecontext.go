package svc

import (
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/api/internal/config"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/rpc/orderclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
