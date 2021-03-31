package service

import (
	"RetailStoreManagementSystem/page/model"
	"RetailStoreManagementSystem/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
)

type UserInfo struct {
	Id int64 `json:"id"`
	RealName string `json:"real_name"`
	UserName string `json:"username"`
	ShopId int64 `json:"shop_id"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}
type UserLoginInp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func Login(c *gin.Context) (*UserInfo,error) {
	inp:=&UserLoginInp{}
	c.BindJSON(inp)
	if inp.Username=="" || inp.Password==""{
		return nil, errors.New("检查参数")
	}
	userInfo:=&model.User{
		UserName: inp.Username,
		Password: inp.Password,
	}
	err:=model.DB().Find(userInfo).Error
	if err!=nil{
		return nil,errors.New("用户名或密码错误")
	}
	res:=&UserInfo{}
	copier.Copy(res,userInfo)
	token,err:=utils.GenerateToken(utils.Struct2Map(*res))
	if err!=nil{
		return nil,errors.New("GenerateToken err")
	}
	err=model.DB().Model(model.User{}).Update(model.User{Token: token}).Where("id=?",userInfo.ID).Error
	if err!=nil{
		return nil,errors.New("set token err")
	}
	res.Token=token
	return res,nil
}
func Logout(c *gin.Context,tokenInfo *TokenInfo)error{
	err:=model.DB().Model(&model.User{}).UpdateColumns(map[string]interface{}{"token":""}).Where("id=?",tokenInfo.Id).Error
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
		log.Warn("请确认参数")
		return errors.New("请确认参数")
	}
	userInfo:=&model.User{
		UserName: username,
	}
	err:=model.DB().Find(userInfo).Error
	if err!=nil && !gorm.IsRecordNotFoundError(err){
		return err
	}
	if userInfo.ID!=0{
		log.Warn("用户已存在")
		return errors.New("用户已存在")
	}
	userInfo=&model.User{
		RealName: realName,
		UserName: username,
		Password: password,
		Phone:    phone,
	}
	err=model.DB().Create(&userInfo).Error
	return err

}
func UpdateUser(c *gin.Context)error{
	userInfo:=&model.User{
		RealName: c.PostForm("real_name"),
		UserName: c.PostForm("username"),
		ShopID:   utils.StringToInt64(c.PostForm("shop_id")),
		Phone:    c.PostForm("phone"),
	}
	return model.DB().Update(userInfo).Error
}
