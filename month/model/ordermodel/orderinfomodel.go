package ordermodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ OrderinfoModel = (*customOrderinfoModel)(nil)

type (
	// OrderinfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderinfoModel.
	OrderinfoModel interface {
		orderinfoModel
	}

	customOrderinfoModel struct {
		*defaultOrderinfoModel
	}
)

// NewOrderinfoModel returns a ordermodel for the database table.
func NewOrderinfoModel(conn sqlx.SqlConn) OrderinfoModel {
	return &customOrderinfoModel{
		defaultOrderinfoModel: newOrderinfoModel(conn),
	}
}
