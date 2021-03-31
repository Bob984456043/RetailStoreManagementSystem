package service

import (
	"RetailStoreManagementSystem/page/model"
	"RetailStoreManagementSystem/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/prometheus/common/log"
	"net/http"
)

type TokenInfo struct {
	Id int64 `json:"id"`
	RealName string `json:"real_name"`
	Username string `json:"username"`
	ShopId int64 `json:"shop_id"`
	Phone string `json:"phone"`
	RoleId int64 `json:"role_id"`
	PermissionList []*model.Permission `json:"permission_list"`
}
func GetTokenInfo(context *gin.Context)(*TokenInfo,error)  {
	token:=context.GetHeader("token")
	if token==""{
		context.JSON(http.StatusForbidden, gin.H{
			"message": "用户认证失败",
		})
		log.Warn("用户认证失败")
		return nil,errors.New("用户认证失败")
	}
	userInfo:=&model.User{
		Token: token,
	}
	err:=model.DB().Find(userInfo).Error
	if err!=nil{
		context.JSON(http.StatusForbidden, gin.H{
			"message": "用户认证失败",
		})
		log.Warn("用户认证失败:%s",err.Error())
		return nil,errors.New("用户认证失败")
	}
	ruleList:=[]*model.Rule{}
	_=model.DB().Find(&ruleList,"role_id=?",userInfo.RoleId).Error
	permissionList:=[]*model.Permission{}
	_=model.DB().Find(&permissionList,"id in (?)",utils.GetInt64Fields(ruleList,"PermissionId")).Error
	res:=&TokenInfo{}
	copier.Copy(res,userInfo)
	res.PermissionList=permissionList
	return res,nil
}
