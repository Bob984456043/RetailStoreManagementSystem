package service

import (
	"RetailStoreManagementSystem/page/model"
	"github.com/gin-gonic/gin"
)

func CreateShop(c *gin.Context,) error {
	shopInfo:=&model.Shop{
		Address:    c.PostForm("address"),
		Name:       c.PostForm("name"),
		Owner:      c.PostForm("owner"),
	}
	err:=model.DB().Create(&shopInfo).Error
	if err!=nil{
		return err
	}
	return nil
}
