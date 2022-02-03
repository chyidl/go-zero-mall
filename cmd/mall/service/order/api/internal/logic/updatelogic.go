package logic

import (
	"context"

	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/api/internal/svc"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/api/internal/types"
	"github.com/chyidl/go-zero-mall/cmd/mall/service/order/rpc/orderclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateLogic {
	return UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.OrderRpc.Update(l.ctx, &orderclient.UpdateRequest{
		Id:     req.Id,
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{}, nil
}
