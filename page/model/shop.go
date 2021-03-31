package model

type Shop struct {
	Address    string    `gorm:"column:address" json:"address"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Owner      string    `gorm:"column:owner" json:"owner"`
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
}

// TableName sets the insert table name for this struct type
func (s *Shop) TableName() string {
	return "shop"
}
