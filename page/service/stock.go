package service

import (
	"RetailStoreManagementSystem/page/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

type GetStockListReq struct {
	name string `json:"name"`
	Barcode string `json:"barcode"`
	sort string `json:"sort"`
	Limit int64 `json:"limit"`
	Offset int64 `json:"offset"`
}
type GetStockDetailResp struct {
	*model.Stock
	Records []*model.StockRecord `json:"records"`
	Sku *model.Sku `json:"sku"`
}
func GetStockList(c *gin.Context,info *TokenInfo)([]*model.Stock,error){
	res:=[]*model.Stock{}
	inp:=&GetStockListReq{}
	c.BindQuery(inp)
	db:=model.DB()
	if inp.name!=""{
		db=db.Where("id in (select id from sku where name like '%?%')",inp.name)
	}
	if inp.Barcode!=""{
		db=db.Where("barcode=?",inp.Barcode)
	}
	if inp.sort!=""{
		db=db.Order(inp.sort)
	}else{
		db=db.Order("id desc")
	}
	err:=model.DB().Find(&res,"shop_id=?",info.ShopId).Limit(inp.Limit).Offset(inp.Offset).Error
	if err!=nil{
		return nil, err
	}
	return res, nil
}

func GetStockDetail(c *gin.Context,info *TokenInfo)(*GetStockDetailResp,error){
	id:=c.Param("id")
	stockInfo:=&model.Stock{}
	err:=model.DB().Find(stockInfo,"id=?",id).Error
	if err!=nil{
		return nil, err
	}
	records:=[]*model.StockRecord{}
	err=model.DB().Find(&records,"sku_id=?",stockInfo.SkuID).Order("create_time desc").Error
	if err!=nil && !gorm.IsRecordNotFoundError(err){
		return nil, err
	}
	skuInfo:=&model.Sku{}
	err=model.DB().Find(skuInfo,"id=?",stockInfo.SkuID).Error
	if err!=nil{
		return nil, err
	}
	res:=&GetStockDetailResp{
		Stock:   stockInfo,
		Records: records,
		Sku: skuInfo,
	}
	return res, nil
}

type WarehousingReq struct {
	Records []*model.StockRecord `json:"records"`
}
func Warehousing(c *gin.Context,info *TokenInfo)error{
	inp:=&WarehousingReq{}
	c.BindJSON(inp)
	db:=model.DB()
	for _,record:=range inp.Records{
		stockInfo:=&model.Stock{}
		err:=db.Find(stockInfo,"sku_id=?",record.SkuID).Error
		if err!=nil{
			return err
		}

		err=db.Create(record).Error
		if err!=nil{
			return err
		}
		err=db.Model(&model.Stock{}).Update(&model.Stock{
			UnitPrice:  record.UnitPrice,
			UpdateTime: time.Now().Unix(),
			CreateTime: record.CreateTime,
			Amount:     record.Amount+stockInfo.Amount,
		}).Error
		if err!=nil{
			return err
		}
	}
	return nil
}
