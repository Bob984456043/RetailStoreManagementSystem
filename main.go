package main

import (
	"RetailStoreManagementSystem/page/service"
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	r:=gin.Default()
	r.POST("/api/user/login", func(context *gin.Context) {
		userInfo,err:=service.Login(context)
		if err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{
				"message":err.Error(),
			})
		}else{
			context.JSON(http.StatusOK,gin.H{
				"message":"登录成功",
				"data":userInfo,
			})
		}
	})
	
	r.POST("/api/user/logout", func(context *gin.Context) {
		tokenInfo,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		err=service.Logout(context,tokenInfo)
		if err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{
				"message":"注销失败",
			})
		}else{
			context.JSON(http.StatusOK,gin.H{
				"message":"注销成功",
			})
		}
	})
	r.POST("/api/user/register", func(context *gin.Context) {
		err:=service.Register(context)
		if err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{
				"message":"注册失败",
			})
		}else{
			context.JSON(http.StatusOK,gin.H{
				"message":"注册成功",
			})
		}
	})
	r.PATCH("/api/user", func(context *gin.Context) {
		_,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		err=service.UpdateUser(context)
		if err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{
				"message":"更新用户信息失败",
			})
		}else{
			context.JSON(http.StatusOK,gin.H{
				"message":"更新用户信息成功",
			})
		}
	})
	r.POST("/api/role", func(context *gin.Context) {
		tokenInfo,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		err=service.CreateRole(context,tokenInfo)
		if err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{
				"message":"创建角色失败",
			})
		}else{
			context.JSON(http.StatusOK,gin.H{
				"message":"创建角色成功",
			})
		}

	})
	r.POST("/api/spu", func(context *gin.Context) {
		tokenInfo,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		err=service.CreateSpu(context,tokenInfo)
		if err!=nil{
			context.JSON(http.StatusBadRequest, err.Error())
		}else {
			context.JSON(http.StatusOK,"创建商品成功")
		}
	})
	r.GET("/api/spu", func(context *gin.Context) {
		tokenInfo,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		res,err:=service.GetSpuList(context,tokenInfo)
		if err!=nil{
			context.JSON(http.StatusBadRequest, err.Error())
		}else {
			context.JSON(http.StatusOK,res)
		}
	})
	r.GET("/api/spu/:id", func(context *gin.Context) {
		tokenInfo,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		res,err:=service.GetSpuDetail(context,tokenInfo)
		if err!=nil{
			context.JSON(http.StatusBadRequest, err.Error())
		}else{
			context.JSON(http.StatusOK,res)
		}

	})
	r.DELETE("/api/spu/:id", func(context *gin.Context) {
		tokenInfo,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		err=service.DelSpu(context,tokenInfo)
		if err!=nil{
			context.JSON(http.StatusBadRequest, err.Error())
		}else{
			context.JSON(http.StatusOK,"删除成功")
		}
	})
	r.DELETE("/api/sku/:id", func(context *gin.Context) {
		tokenInfo,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		err=service.DelSku(context,tokenInfo)
		if err!=nil{
			context.JSON(http.StatusBadRequest, err.Error())
		}else{
			context.JSON(http.StatusOK,"删除成功")
		}
	})
	r.POST("/api/sku", func(context *gin.Context) {
		tokenInfo,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		err=service.AddSku(context,tokenInfo)
		if err!=nil{
			context.JSON(http.StatusBadRequest, err.Error())
		}else{
			context.JSON(http.StatusOK,"创建成功")
		}
	})
	r.GET("/api/stock", func(context *gin.Context) {
		tokenInfo,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		res,err:=service.GetStockList(context,tokenInfo)
		if err!=nil{
			context.JSON(http.StatusBadRequest, err.Error())
		}else{
			context.JSON(http.StatusOK,res)
		}
	})
	r.GET("/api/stock/:id", func(context *gin.Context) {
		tokenInfo,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		res,err:=service.GetStockDetail(context,tokenInfo)
		if err!=nil{
			context.JSON(http.StatusBadRequest, err.Error())
		}else{
			context.JSON(http.StatusOK,res)
		}
	})
	r.POST("/api/stock/warehousing", func(context *gin.Context) {
		tokenInfo,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		err=service.Warehousing(context,tokenInfo)
		if err!=nil{
			context.JSON(http.StatusBadRequest, err.Error())
		}else{
			context.JSON(http.StatusOK,"入库成功")
		}
	})
	r.Run(":8080")
}
