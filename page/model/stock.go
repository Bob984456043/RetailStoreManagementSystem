package model

import "github.com/guregu/null"

type Stock struct {
	ShopID     int64     `gorm:"column:shop_id" json:"shop_id"`
	Barcode    string    `gorm:"column:barcode" json:"barcode"`
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	SaleNum    null.Int  `gorm:"column:sale_num" json:"sale_num"`
	SkuID      int64     `gorm:"column:sku_id" json:"sku_id"`
	UpdateTime null.Time `gorm:"column:update_time" json:"update_time"`
	AddUpNum   null.Int  `gorm:"column:add_up_num" json:"add_up_num"`
	CreateTime null.Time `gorm:"column:create_time" json:"create_time"`
	Num        null.Int  `gorm:"column:num" json:"num"`
}

// TableName sets the insert table name for this struct type
func (s *Stock) TableName() string {
	return "stock"
}
