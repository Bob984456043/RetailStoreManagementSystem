package model

import (
	"time"
)

type User struct {
	RealName   string      `gorm:"column:real_name" json:"real_name"`
	ShopID     int64       `gorm:"column:shop_id" json:"shop_id"`
	UpdateTime int64   `gorm:"column:update_time" json:"update_time"`
	CreateTime int64   `gorm:"column:create_time" json:"create_time"`
	Password   string      `gorm:"column:password" json:"password"`
	Token      string `gorm:"column:token" json:"token"`
	UserName   string      `gorm:"column:user_name" json:"user_name"`
	ID         int64       `gorm:"column:id;primary_key" json:"id"`
	Phone      string      `gorm:"column:phone" json:"phone"`
	RoleId int64 `gorm:column:"role_id" json:"role_id"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate() (err error) {
	u.CreateTime=time.Now().Unix()
	return
}
