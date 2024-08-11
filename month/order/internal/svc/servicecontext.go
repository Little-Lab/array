package svc

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"order/internal/config"
	"zg6/2112a-6/month/model/inventorymodel"
	"zg6/2112a-6/month/model/ordermodel"
)

type ServiceContext struct {
	Config         config.Config
	DB             *sql.DB
	OrderModel     ordermodel.OrderinfoModel
	InventoryModel inventorymodel.GoodsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := sqlx.NewMysql(c.DB.DataSource).RawDB()
	return &ServiceContext{
		Config:         c,
		DB:             db,
		OrderModel:     ordermodel.NewOrderinfoModel(sqlx.NewMysql(c.DB.DataSource)),
		InventoryModel: inventorymodel.NewGoodsModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
