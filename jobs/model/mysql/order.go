package mysql

import (
	"time"
)

type OrderInfo struct {
	Id         int       `json:"id"`
	OrderSn    string    `json:"order_sn"`
	GoodsId    int       `json:"goods_id"`
	UserId     int       `json:"user_id"`
	PayType    int       `json:"pay_type"`
	Status     int       `json:"status"`
	Count      int       `json:"count"`
	OrderMount float64   `json:"order_mount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	IsDeleted  int       `json:"is_deleted"`
}
