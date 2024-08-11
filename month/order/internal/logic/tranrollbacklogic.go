package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"
	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranRollbackLogic {
	return &TranRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranRollbackLogic) TranRollback(in *order.TranDelOrderRequest) (*order.Response, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("删除订单  , in: %+v \n", in)

	orderinfo, err := l.svcCtx.OrderModel.FindLastOneByUserIdGoodsId(in.Id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if orderinfo != nil {
		barrier, _ := dtmgrpc.BarrierFromGrpc(l.ctx)

		err = barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {

			err = l.svcCtx.OrderModel.Update(tx, orderinfo.Id)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, errors.New(err.Error())
		}
	}
	return &order.Response{
		Success: true,
	}, nil
}
