package logic

import (
	"context"
	"order/order"

	"bff/internal/svc"
	"bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateOrderRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.OrderSrv.CreateOrder(l.ctx, &order.TranCreateOrderRequest{
		GoodsId: req.GoodsId,
		UserId:  req.UserId,
		Num:     req.Num,
		PayType: req.PayType,
	})
	if err != nil {
		return &types.Response{
			Code: 400,
			Msg:  err.Error(),
		}, nil
	}
	return &types.Response{
		Code: 200,
		Msg:  "success",
	}, nil
}
