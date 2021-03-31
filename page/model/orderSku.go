package model

type OrderSku struct {
	SkuID      int64     `gorm:"column:sku_id" json:"sku_id"`
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	OrderID    int64     `gorm:"column:order_id" json:"order_id"`
	SalePrice  float64   `gorm:"column:sale_price" json:"sale_price"`
}

// TableName sets the insert table name for this struct type
func (o *OrderSku) TableName() string {
	return "order_sku"
}
