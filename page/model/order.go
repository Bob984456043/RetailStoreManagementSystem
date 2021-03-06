package model

type Order struct {
	Amount     int       `gorm:"column:amount" json:"amount"`
	Count      float64   `gorm:"column:count" json:"count"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	Opreator   int64     `gorm:"column:opreator" json:"opreator"`
	ShopID     int64     `gorm:"column:shop_id" json:"shop_id"`
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (o *Order) TableName() string {
	return "order"
}
