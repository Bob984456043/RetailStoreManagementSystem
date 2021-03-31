package model

import (
	"context"
	"github.com/guregu/null"
)

type User struct {
	RealName   string      `gorm:"column:real_name" json:"real_name"`
	ShopID     int64       `gorm:"column:shop_id" json:"shop_id"`
	UpdataTime null.Time   `gorm:"column:updata_time" json:"updata_time"`
	CreateTime null.Time   `gorm:"column:create_time" json:"create_time"`
	Password   string      `gorm:"column:password" json:"password"`
	Token      null.String `gorm:"column:token" json:"token"`
	UserName   string      `gorm:"column:user_name" json:"user_name"`
	ID         int64       `gorm:"column:id;primary_key" json:"id"`
	Phone      string      `gorm:"column:phone" json:"phone"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}



func (u *User) GetUserInfoByToken(ctx context.Context,token string)(*User,error)  {
	return nil,nil
}
func (u *User)GetUserInfoByUsernameAndPassword(ctx context.Context,username string,password string)(*User,error){
	return nil, nil
}
func (u *User)GetUserInfoByUsername(ctx context.Context,username string)(*User,error){
	return nil, nil
}
func (u *User)SetToken(ctx context.Context,token string)error{
	return nil
}
func (u *User)Create(ctx context.Context,user *User)error{
	return nil
}
func(u *User)Update(ctx context.Context,user *User)error{
	return nil
}
