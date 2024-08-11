package logic

import (
	"context"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"order/order"
	"stock/stock"

	"bff/internal/svc"
	"bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BffLogic {
	return &BffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var dtmServer = "etcd://localhost:2379/dtmservice"

func (l *BffLogic) Bff(req *types.BFFRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	orderRpcBusiServer, err := l.svcCtx.Config.OrderSrv.BuildTarget()
	if err != nil {
		return nil, fmt.Errorf("下单异常超时")
	}
	stockRpcBusiServer, err := l.svcCtx.Config.StockSrv.BuildTarget()
	if err != nil {
		return nil, fmt.Errorf("库存异常超时")
	}

	gid := dtmgrpc.MustGenGid(dtmServer)
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid)
	saga.Add(orderRpcBusiServer+"/order.Order/TranCreateOrder", orderRpcBusiServer+"/order.Order/TranRollback", &order.TranCreateOrderRequest{
		GoodsId: req.GoodsId,
		UserId:  req.UserId,
		Num:     req.Num,
		PayType: req.PayType,
	})
	saga.Add(stockRpcBusiServer+"/stock.Stock/TranUpdateInventory", stockRpcBusiServer+"/stock.Stock/TranUpdateInventoryRollback", &stock.TranUpdateInventoryRequest{
		GoodsId: req.GoodsId,
		Stock:   req.Num,
	})

	err = saga.Submit()
	if err != nil {
		return nil, fmt.Errorf("submit data to  dtm-server err  : %+v \n", err)
	}
	return &types.Response{
		Code: 200,
		Msg:  "success",
	}, nil
}
