package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dtm-labs/client/dtmgrpc"

	"stock/internal/svc"
	"stock/stock"

	"github.com/zeromicro/go-zero/core/logx"
)

type TranUpdateInventoryRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranUpdateInventoryRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranUpdateInventoryRollbackLogic {
	return &TranUpdateInventoryRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranUpdateInventoryRollbackLogic) TranUpdateInventoryRollback(in *stock.TranUpdateInventoryRequest) (*stock.Response, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("库存回滚 in : %+v \n", in)

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	err = barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {
		err = l.svcCtx.InventoryModel.AddStock(tx, in.GoodsId, in.Stock)
		if err != nil {
			return fmt.Errorf("添加库存失败 err : %v ,goodsId:%d , num :%d", err, in.GoodsId, in.Stock)
		}
		return nil
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &stock.Response{
		Success: true,
	}, nil
}
