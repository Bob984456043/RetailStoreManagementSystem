package model

import "github.com/guregu/null"

type Spu struct {
	CategoryID    int64       `gorm:"column:category_id" json:"category_id"`
	MemberPrice   null.Float  `gorm:"column:member_price" json:"member_price"`
	Name          string      `gorm:"column:name" json:"name"`
	OriginalPrice float64     `gorm:"column:original_price" json:"original_price"`
	ShopID        int64       `gorm:"column:shop_id" json:"shop_id"`
	UpdateTime    null.Time   `gorm:"column:update_time" json:"update_time"`
	CreateTime    null.Time   `gorm:"column:create_time" json:"create_time"`
	ID            int64       `gorm:"column:id;primary_key" json:"id"`
	ImgURL        null.String `gorm:"column:img_url" json:"img_url"`
	SpecValues    string      `gorm:"column:spec_values" json:"spec_values"`
}

// TableName sets the insert table name for this struct type
func (s *Spu) TableName() string {
	return "spu"
}

