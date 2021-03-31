package model

type Role struct {
	UpdateTime  int64 `gorm:"column:update_time" json:"update_time"`
	CreateTime  int64 `gorm:"column:create_time" json:"create_time"`
	ID          int64     `gorm:"column:id;primary_key" json:"id"`
	IsCustomer int32  `gorm:"column:is_customer" json:"is_customer"`
	Name        string    `gorm:"column:name" json:"name"`
	ShopID      int64     `gorm:"column:shop_id" json:"shop_id"`
}

// TableName sets the insert table name for this struct type
func (r *Role) TableName() string {
	return "role"
}
