package model

type Permission struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	UpdateTime int64 `gorm:"column:update_time" json:"update_time"`
	URL        string    `gorm:"column:url" json:"url"`
	CreateTime int64 `gorm:"column:create_time" json:"create_time"`
}

// TableName sets the insert table name for this struct type
func (p *Permission) TableName() string {
	return "permission"
}
