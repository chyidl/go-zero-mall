package svc

import (
	"github.com/chyidl/go-zero-mall/cmd/mall/service/product/api/internal/config"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/product/rpc/productclient"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
