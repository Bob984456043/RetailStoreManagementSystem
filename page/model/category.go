package model

type Category struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	Level      string    `gorm:"column:level" json:"level"`
	Name       string    `gorm:"column:name" json:"name"`
	ParentID   int64     `gorm:"column:parent_id" json:"parent_id"`
	ShopID     int64     `gorm:"column:shop_id" json:"shop_id"`
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
}

// TableName sets the insert table name for this struct type
func (c *Category) TableName() string {
	return "category"
}
