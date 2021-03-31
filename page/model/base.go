package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
)

func DB()*gorm.DB{
	db, err := gorm.Open("mysql", "root:qwe984456043@/retail_store_management?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		log.Warn("get db err")
		return nil
	}
	db.LogMode(true)
	return 	db
}
