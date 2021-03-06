package main

import (
	"flag"
	"fmt"

	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/rpc/internal/config"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/rpc/internal/server"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/rpc/internal/svc"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/rpc/order"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewOrderServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		order.RegisterOrderServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
