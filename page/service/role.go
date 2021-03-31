package service

import (
	"RetailStoreManagementSystem/page/model"
	"RetailStoreManagementSystem/utils"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type Role struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	PermissionList []*model.Permission `json:"permission_list"`
	ShopId int64 `json:"shop_id"`
	IsConsumer int32 `json:"is_consumer"`
	
}
func DeleteRole(c *gin.Context,info *TokenInfo)error{
	id:=c.PostForm("id")
	if id==""{
		return errors.New("请确认参数")
	}
	roleInfo:=&model.Role{
		ID: utils.ToInt64(id),
	}
	err:=model.DB().Find(roleInfo).Error
	if err!=nil{
		return err
	}
	if roleInfo.ShopID!=info.ShopId{
		return errors.New("diff shop")
	}
	db:=model.DB()
	db.Begin()
	err=db.Delete(&model.Role{
		ID: utils.ToInt64(id),
	}).Error
	if err!=nil{
		db.Rollback()
		return err
	}
	err=db.Delete(&model.Rule{
		RoleId:       utils.ToInt64(id),
	}).Error
	if err!=nil{
		db.Rollback()
		return err
	}
	db.Commit()
	return err
}
func CreateRole(c *gin.Context,info *TokenInfo)error{
	name:=c.PostForm("name")
	isConsumer:=c.PostForm("is_consumer")
	if name=="" || isConsumer==""{
		return errors.New("请确认参数")
	}
	return model.DB().Create(&model.Role{
		IsCustomer: int32(utils.ToInt64(isConsumer)),
		Name:       name,
		ShopID:     info.ShopId,
	}).Error
}
func GetRoleList(c *gin.Context,info *TokenInfo)([]*Role,error){
	name:=c.Query("name")
	isConsumer:=c.Query("is_consumer")
	res:=[]*Role{}
	err:=model.DB().Find(&res,"name like '%?%' and is_consumer=? and shop_id=?",name,int32(utils.ToInt64(isConsumer)),info.ShopId).Error
	if err!=nil{
		return nil,err
	}
	return res,nil
}
func GetRoleDetail(c *gin.Context,info *TokenInfo)(*Role,error){
	roleId:=c.Param("id")
	roleInfo:=&model.Role{
		ID: utils.ToInt64(roleId),
	}
	err:=model.DB().Find(roleInfo).Error
	if err!=nil{
		return nil,err
	}
	if roleInfo.ShopID!=info.ShopId{
		return nil,errors.New("diff shop")
	}
	res:=&Role{}
	_=copier.Copy(res,roleInfo)
	if roleInfo.IsCustomer==0{
		ruleList:=make([]*model.Rule,0)
		err:=model.DB().Find(&ruleList,"role_id=?",utils.ToInt64(roleId)).Error
		if err!=nil{
			return nil,err
		}
		permissionList:=make([]*model.Permission,0)
		err=model.DB().Find(&permissionList,"id in (?)",utils.GetInt64Fields(permissionList,"ID")).Error
		if err!=nil{
			return nil,err
		}
		res.PermissionList=permissionList
	}
	return res,nil
	
}
func AddPermission(c *gin.Context,info *TokenInfo)error{
	permissionIdsString:=c.PostForm("permission_ids")
	roleId:=c.PostForm("role_id")
	roleInfo:=&model.Role{}
	err:=model.DB().Find(roleInfo,"id=?",utils.ToInt64(roleId)).Error
	if err!=nil{
		return err
	}
	if roleInfo.ShopID!=info.ShopId{
		return errors.New("diff shop")
	}
	permissionIds:=make([]int64,0)
	_=json.Unmarshal([]byte(permissionIdsString),&permissionIds)
	db:=model.DB()
	db.Begin()
	for _,p:=range permissionIds{
		err:=db.Create(&model.Rule{
			RoleId:       utils.ToInt64(roleId),
			PermissionId: p,
		}).Error
		if err!=nil{
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return err
}
func RemovePermission(c *gin.Context,info *TokenInfo)error{
	permissionIdsString:=c.PostForm("permission_ids")
	roleId:=c.PostForm("role_id")
	roleInfo:=&model.Role{}
	err:=model.DB().Find(roleInfo,"id=?",utils.ToInt64(roleId)).Error
	if err!=nil{
		return err
	}
	if roleInfo.ShopID!=info.ShopId{
		return errors.New("diff shop")
	}
	permissionIds:=make([]int64,0)
	_=json.Unmarshal([]byte(permissionIdsString),&permissionIds)
	db:=model.DB()
	db.Begin()
	for _,p:=range permissionIds{
		err:=db.Delete(&model.Rule{
			RoleId:       utils.ToInt64(roleId),
			PermissionId: p,
		}).Error
		if err!=nil{
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return err
}
func GetPermissionList(c *gin.Context,info *TokenInfo)([]*model.Permission,error){
	permissionList:=make([]*model.Permission,0)
	err:=model.DB().Find(&permissionList).Error
	if err!=nil{
		return nil, err
	}
	return permissionList,nil

}
