package model

type Stock struct {
	ShopID     int64     `gorm:"column:shop_id" json:"shop_id"`
	Barcode    string    `gorm:"column:barcode" json:"barcode"`
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	UnitPrice    int64  `gorm:"column:unit_price" json:"unit_price"`
	SkuID      int64     `gorm:"column:sku_id" json:"sku_id"`
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
	Amount        int64  `gorm:"column:amount" json:"amount"`
}

// TableName sets the insert table name for this struct type
func (s *Stock) TableName() string {
	return "stock"
}
