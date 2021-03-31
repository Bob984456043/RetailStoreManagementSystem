package service

import (
	"RetailStoreManagementSystem/page/model"
	"RetailStoreManagementSystem/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserInfo struct {
	Id int64
	RealName string
	Username string
	ShopId int64
	Phone string
	Token string
}
func Login(c *gin.Context) (*UserInfo,error) {
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	if username=="" || password==""{
		return nil, errors.New("用户名或密码错误")
	}
	userInfo,err:=(&model.User{}).GetUserInfoByUsernameAndPassword(c.Request.Context(),username,password)
	if err!=nil{
		return nil,errors.New("用户名或密码错误")
	}
	res:=&UserInfo{}
	copier.Copy(res,userInfo)
	token,err:=utils.GenerateToken(utils.Struct2Map(res))
	if err!=nil{
		return nil,errors.New("GenerateToken err")
	}
	err=(&model.User{}).SetToken(c.Request.Context(),token)
	if err!=nil{
		return nil,errors.New("set token err")
	}
	res.Token=token
	return res,nil
}
func Logout(c *gin.Context)error{
	username:=c.PostForm("username")
	if username==""{
		return errors.New("username is needed")
	}
	err:=(&model.User{}).SetToken(c.Request.Context(),"")
	if err!=nil{
		return errors.New("setToken err")
	}
	return nil
}
func Register(c *gin.Context)error{
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	realName:=c.PostForm("real_name")
	phone:=c.PostForm("phone")
	if username==""||password==""||realName==""||phone==""{
		return errors.New("请确认参数")
	}
	userInfo,err:=(&model.User{}).GetUserInfoByUsername(c.Request.Context(),username)
	if err!=nil{
		return err
	}
	if userInfo!=nil || userInfo.Id!=0{
		return errors.New("用户已存在")
	}
	err=(&model.User{}).Create(c.Request.Context(),&model.User{
		RealName: realName,
		Username: username,
		Password: password,
		Phone:    phone,
	})
	return err

}
func Update(c *gin.Context)error{
	userInfo:=&model.User{
		RealName: c.PostForm("real_name"),
		Username: c.PostForm("username"),
		ShopId:   utils.StringToInt64(c.PostForm("shop_id")),
		Phone:    c.PostForm("phone"),
	}
	return (&model.User{}).Update(c.Request.Context(),userInfo)
}
