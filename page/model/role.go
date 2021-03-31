package model

import (
	"github.com/guregu/null"
)

type Role struct {
	UpdateTime  null.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime  null.Time `gorm:"column:create_time" json:"create_time"`
	ID          int64     `gorm:"column:id;primary_key" json:"id"`
	Is_customer null.Int  `gorm:"column:is_ customer" json:"is_ customer"`
	Name        string    `gorm:"column:name" json:"name"`
	ShopID      int64     `gorm:"column:shop_id" json:"shop_id"`
}

// TableName sets the insert table name for this struct type
func (r *Role) TableName() string {
	return "role"
}
