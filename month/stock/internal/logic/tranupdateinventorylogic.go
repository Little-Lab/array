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

type TranUpdateInventoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranUpdateInventoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranUpdateInventoryLogic {
	return &TranUpdateInventoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranUpdateInventoryLogic) TranUpdateInventory(in *stock.TranUpdateInventoryRequest) (*stock.Response, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("扣库存start....")

	goods, err := l.svcCtx.InventoryModel.FindOneByGoodsId(in.GoodsId)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if goods == nil || goods.Stock < in.Stock {
		return nil, errors.New(err.Error())
	}

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	err = barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {
		sqlResult, _ := l.svcCtx.InventoryModel.DecuctStock(tx, in.GoodsId, in.Stock)

		affected, _ := sqlResult.RowsAffected()

		if affected <= 0 {
			return errors.New(err.Error())
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
