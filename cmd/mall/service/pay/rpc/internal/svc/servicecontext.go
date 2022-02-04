package svc

import (
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/rpc/orderclient"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/pay/model"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/pay/rpc/internal/config"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/user/rpc/userclient"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PayModel model.PayModel

	UserRpc  userclient.User
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
