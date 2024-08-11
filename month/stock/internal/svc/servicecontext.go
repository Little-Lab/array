package svc

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"stock/internal/config"
	"zg6/2112a-6/month/model/inventorymodel"
)

type ServiceContext struct {
	Config         config.Config
	DB             *sql.DB
	InventoryModel inventorymodel.GoodsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := sqlx.NewMysql(c.DB.DataSource).RawDB()
	return &ServiceContext{
		Config:         c,
		DB:             db,
		InventoryModel: inventorymodel.NewGoodsModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
