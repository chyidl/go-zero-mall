package svc

import (
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/model"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/rpc/internal/config"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/product/rpc/productclient"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/user/rpc/userclient"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderModel model.OrderModel

	UserRpc    userclient.User
	ProductRpc productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
