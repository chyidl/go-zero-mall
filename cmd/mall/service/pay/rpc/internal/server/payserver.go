// Code generated by goctl. DO NOT EDIT!
// Source: pay.proto

package server

import (
	"context"

	"github.com/chyidl/go-zero-mall/cmd/mall/service/pay/rpc/internal/logic"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/pay/rpc/internal/svc"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/pay/rpc/pay"
)

type PayServer struct {
	svcCtx *svc.ServiceContext
}

func NewPayServer(svcCtx *svc.ServiceContext) *PayServer {
	return &PayServer{
		svcCtx: svcCtx,
	}
}

func (s *PayServer) Create(ctx context.Context, in *pay.CreateRequest) (*pay.CreateResponse, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *PayServer) Detail(ctx context.Context, in *pay.DetailRequest) (*pay.DetailResponse, error) {
	l := logic.NewDetailLogic(ctx, s.svcCtx)
	return l.Detail(in)
}

func (s *PayServer) Callback(ctx context.Context, in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	l := logic.NewCallbackLogic(ctx, s.svcCtx)
	return l.Callback(in)
}
