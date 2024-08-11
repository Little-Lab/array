package mysql

import (
	"time"
	"zg6/2112a-6/jobs/global"
)

type Goods struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	OnSale     int       `json:"on_sale"`
	GoodsBrief string    `json:"goods_brief"`
	Price      float64   `json:"price"`
	UserId     int       `json:"user_id"`
	IsHot      int       `json:"is_hot"`
	Types      string    `json:"types"`
	Stock      int       `json:"stock"`
	Images     string    `json:"images"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	IsDeleted  int       `json:"is_deleted"`
}

func UpdateStock(id, stock int) error {
	return global.DB.Exec("update goods set stock = stock + ? where id = ?", stock, id).Error
}
