package logic

import (
	"context"
	"errors"
	"fmt"
	"order/utils"
	"time"
	"zg6/2112a-6/month/model/ordermodel"

	"order/internal/svc"
	"order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.TranCreateOrderRequest) (*order.Response, error) {
	// todo: add your logic here and delete this line
	//计算金额
	goods, err := l.svcCtx.InventoryModel.FindOneByGoodsId(in.GoodsId)
	if err != nil {
		return nil, errors.New(err.Error())
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
	_, err = l.svcCtx.OrderModel.Inserts(l.ctx, orderinfo)
	if err != nil {
		return nil, fmt.Errorf("创建订单失败 err : %v , order:%+v \n", err, orderinfo)
	}

	result, _ := l.svcCtx.InventoryModel.DecuctStocks(l.ctx, in.GoodsId, in.Num)

	affected, _ := result.RowsAffected()

	if affected < 1 {
		return nil, errors.New("库存扣减失败")
	}
	return &order.Response{
		Success: true,
	}, nil
}
