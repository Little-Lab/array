package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"order/utils"
	"time"
	"zg6/2112a-6/month/model/ordermodel"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"order/internal/svc"
	"order/order"
)

type TranCreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTranCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TranCreateOrderLogic {
	return &TranCreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TranCreateOrderLogic) TranCreateOrder(in *order.TranCreateOrderRequest) (*order.Response, error) {
	// todo: add your logic here and delete this line
	fmt.Printf("创建订单 in : %+v \n", in)

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	if err = barrier.CallWithDB(l.svcCtx.DB, func(tx *sql.Tx) error {

		//计算金额
		goods, err := l.svcCtx.InventoryModel.FindOneByGoodsId(in.GoodsId)
		if err != nil {
			return err
		}
		mount := goods.Price * float64(in.Num)

		orderinfo := new(ordermodel.Orderinfo)
		orderinfo.GoodsId = in.GoodsId
		orderinfo.Count = in.Num
		orderinfo.UserId = in.UserId
		orderinfo.PayType = in.PayType
		orderinfo.CreatedAt = time.Now()
		orderinfo.UpdatedAt = time.Now()
		orderinfo.OrderSn = utils.GenerateSN()
		orderinfo.OrderMount = mount
		_, err = l.svcCtx.OrderModel.Insert(tx, orderinfo)
		if err != nil {
			return fmt.Errorf("创建订单失败 err : %v , order:%+v \n", err, orderinfo)
		}
		return nil
	}); err != nil {
		return nil, errors.New(err.Error())
	}

	return &order.Response{
		Success: true,
	}, nil
}
