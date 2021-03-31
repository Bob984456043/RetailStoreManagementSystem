package service

import (
	"RetailStoreManagementSystem/page/model"
	"github.com/gin-gonic/gin"
)

func AddSku(c *gin.Context,info *TokenInfo) error {
	skuInfo:=&model.Sku{}
	c.BindJSON(skuInfo)
	db:=model.DB()
	err:=db.Create(skuInfo).Error
	if err!=nil{
		return err
	}
	err=db.Create(&model.Stock{
		ShopID:     skuInfo.ShopID,
		Barcode:    skuInfo.Barcode,
		SkuID:      skuInfo.ID,
	}).Error
	if err!=nil{
		return err
	}
	return nil
}
func DelSku(c *gin.Context,info *TokenInfo)error{
	id:=c.Param("id")
	err:=model.DB().Delete(&model.Sku{},"id=?",id).Error
	if err!=nil{
		return err
	}
	return nil
}
