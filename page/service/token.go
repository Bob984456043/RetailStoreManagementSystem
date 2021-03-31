package service
import (
	"RetailStoreManagementSystem/page/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/prometheus/common/log"
	"net/http"
)

type TokenInfo struct {
	Id int64
	RealName string
	Username string
	ShopId int64
	Phone string
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
	userInfo,err:=(&model.User{}).GetUserInfoByToken(context.Request.Context(),token)
	if err!=nil{
		context.JSON(http.StatusForbidden, gin.H{
			"message": "用户认证失败",
		})
		log.Warn("用户认证失败:%s",err.Error())
		return nil,errors.New("用户认证失败")
	}
	if userInfo.Id==0{
		context.JSON(http.StatusForbidden, gin.H{
			"message": "用户认证失败",
		})
		log.Warn("用户认证失败:id 为空")
		return nil,errors.New("用户认证失败")
	}
	res:=&TokenInfo{}
	copier.Copy(res,userInfo)
	return res,nil
}
