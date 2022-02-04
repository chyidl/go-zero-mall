package svc

import (
	"github.com/chyidl/go-zero-mall/cmd/mall/service/pay/api/internal/config"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/pay/rpc/payclient"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PayRpc payclient.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
