package logic

import (
	"context"

	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/api/internal/svc"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/api/internal/types"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/rpc/orderclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveLogic {
	return RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.OrderRpc.Remove(l.ctx, &orderclient.RemoveRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.RemoveResponse{}, nil
}