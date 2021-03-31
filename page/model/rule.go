package model

import "github.com/guregu/null"

type Rule struct {
	UpdateTime null.Time `gorm:"column:update_time" json:"update_time"`
	V0         string    `gorm:"column:v0" json:"v0"`
	V1         string    `gorm:"column:v1" json:"v1"`
	CreateTime null.Time `gorm:"column:create_time" json:"create_time"`
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
}

// TableName sets the insert table name for this struct type
func (r *Rule) TableName() string {
	return "rule"
}

