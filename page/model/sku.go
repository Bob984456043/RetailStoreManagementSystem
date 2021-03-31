package model

import "github.com/guregu/null"

type Sku struct {
	ImgURL        null.String `gorm:"column:img_url" json:"img_url"`
	Name          string      `gorm:"column:name" json:"name"`
	OriginalPrice float64     `gorm:"column:original_price" json:"original_price"`
	UpdateTime    null.Time   `gorm:"column:update_time" json:"update_time"`
	ID            int64       `gorm:"column:id;primary_key" json:"id"`
	CategoryID    int64       `gorm:"column:category_id" json:"category_id"`
	CreateTime    null.Time   `gorm:"column:create_time" json:"create_time"`
	MemberPrice   float64     `gorm:"column:member_price" json:"member_price"`
	ShopID        int64       `gorm:"column:shop_id" json:"shop_id"`
	Barcode       string      `gorm:"column:barcode" json:"barcode"`
}

// TableName sets the insert table name for this struct type
func (s *Sku) TableName() string {
	return "sku"
}

