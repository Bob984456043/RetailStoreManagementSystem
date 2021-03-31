package service

import (
	"RetailStoreManagementSystem/page/model"
	"github.com/gin-gonic/gin"
)

func GetCategoryList(c *gin.Context,info *TokenInfo)([]*model.Category,error){
	res:=[]*model.Category{}
	err:=model.DB().Find(&res).Error
	if err!=nil{
		return nil, err
	}
	return res,nil

}
