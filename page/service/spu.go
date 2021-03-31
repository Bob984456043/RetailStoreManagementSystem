package service

import (
	"RetailStoreManagementSystem/page/model"
	"RetailStoreManagementSystem/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)
type Spu struct {
	CategoryID    int64       `gorm:"column:category_id" json:"category_id"`
	MemberPrice   int64  `gorm:"column:member_price" json:"member_price"`
	Name          string      `gorm:"column:name" json:"name"`
	OriginalPrice int64     `gorm:"column:original_price" json:"original_price"`
	ShopID        int64       `gorm:"column:shop_id" json:"shop_id"`
	UpdateTime    int64   `gorm:"column:update_time" json:"update_time"`
	CreateTime    int64   `gorm:"column:create_time" json:"create_time"`
	ID            int64       `gorm:"column:id;primary_key" json:"id"`
	ImgURL        string `gorm:"column:img_url" json:"img_url"`
	SpecValues    string      `gorm:"column:spec_values" json:"spec_values"`
	Skus []*model.Sku `json:"skus"`
}
func CreateSpu(c *gin.Context,info *TokenInfo)error{

	spu:=&Spu{}
	spuInfo:=&model.Spu{}
	c.BindJSON(spu)
	fmt.Println(spu.Skus[0])
	_=copier.Copy(spuInfo,spu)
	db:=model.DB()
	err:=db.Create(spuInfo).Error
	if err!=nil{
		return err
	}
	for _,sku:=range spu.Skus{
		err:=db.Create(sku).Error
		if err!=nil{
			return err
		}
		err=db.Create(&model.Stock{
			ShopID:     sku.ShopID,
			Barcode:    sku.Barcode,
			SkuID:      sku.ID,
		}).Error
		if err!=nil{
			return err
		}
	}
	return nil
}

type GetSpuListReq struct {
	Name string `json:"name"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
	CategoryId int64 `json:"category_id"`
}
func GetSpuList(c *gin.Context,info *TokenInfo)([]*model.Spu,error){
	inp:=&GetSpuListReq{}
	c.BindQuery(&inp)
	res:=[]*model.Spu{}
	db:=model.DB()
	if inp.Name!=""{
		db=db.Where("name like '%?%'",inp.Name)
	}
	if inp.CategoryId!=0{
		db=db.Where("category_id=?",inp.CategoryId)
	}
	err:=db.Find(&res).Limit(utils.ToInt64(inp.Limit)).Offset(utils.ToInt64(inp.Offset)).Error
	if err!=nil{
		return nil, err
	}
	return res,nil
}
func GetSpuDetail(c *gin.Context,info *TokenInfo)(*Spu,error){
	id:=c.Param("id")
	spuinfo:=&model.Spu{
		ID: utils.ToInt64(id),
	}
	err:=model.DB().Find(spuinfo).Error
	if err!=nil{
		return nil, err
	}
	skus:=[]*model.Sku{}
	err=model.DB().Find(&skus,"spu_id=?",spuinfo.ID).Error
	if err!=nil{
		return nil,err
	}
	res:=&Spu{}
	_=copier.Copy(res,spuinfo)
	res.Skus=skus
	return res,nil
}
func DelSpu(c *gin.Context,info *TokenInfo)error{
	id:=c.Param("id")
	err:=model.DB().Delete(&model.Spu{},"id=?",id).Error
	if err!=nil{
		return err
	}
	skus:=[]*model.Sku{}
	err=model.DB().Find(&skus,"spu_id=?",id).Error
	if err!=nil{
		return err
	}
	err=model.DB().Delete(&model.Sku{},"spu_id=?",id).Error
	if err!=nil{
		return err
	}
	err=model.DB().Delete(&model.Stock{},"sku_id in (?)",utils.GetInt64Fields(skus,"ID")).Error
	if err!=nil{
		return err
	}
	return nil
}
