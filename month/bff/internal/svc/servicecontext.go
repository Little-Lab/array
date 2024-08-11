package svc

import (
	"bff/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"order/order"
	"order/orderclient"
	"stock/stock"
	"stock/stockclient"
)

type ServiceContext struct {
	Config   config.Config
	StockSrv stock.StockClient
	OrderSrv order.OrderClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		StockSrv: stockclient.NewStock(zrpc.MustNewClient(c.StockSrv)),
		OrderSrv: orderclient.NewOrder(zrpc.MustNewClient(c.OrderSrv)),
	}
}
