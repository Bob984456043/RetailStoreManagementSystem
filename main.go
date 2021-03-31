package main

import (
	"RetailStoreManagementSystem/page/service"
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	r:=gin.Default()
	
	r.POST("/user/login", func(context *gin.Context) {
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
	
	r.POST("/user/logout", func(context *gin.Context) {
		_,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		err=service.Logout(context)
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
	r.POST("/user/register", func(context *gin.Context) {
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
	r.POST("/user/update", func(context *gin.Context) {
		_,err:=service.GetTokenInfo(context)
		if err!=nil{
			return
		}
		err=service.Update(context)
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
	
	r.Run(":8080")
}
