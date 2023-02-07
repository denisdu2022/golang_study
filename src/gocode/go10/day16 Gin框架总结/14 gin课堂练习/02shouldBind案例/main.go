package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 结构体
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 实例化结构体全局对象
var login Login

// jsonHandler
func jsonHandler(ctx *gin.Context) {
	//ShouldBind()会根据Content-Type自行选择绑定器
	if err := ctx.ShouldBind(&login); err == nil {
		fmt.Printf("%#v\n", login)
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"user":     login.User,
			"password": login.Password,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

// formHandler
func formHandler(ctx *gin.Context) {
	//ShouldBind()会根据content-type类型自行选择绑定器
	if err := ctx.ShouldBind(&login); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"user":     login.User,
			"password": login.Password,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

// querystringHandler
func querystringHandle(ctx *gin.Context) {
	//ShouldBind()会根据Content-Type自行选择绑定器
	if err := ctx.ShouldBind(&login); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"user":     login.User,
			"password": login.Password,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func main() {

	//获取路由对象
	r := gin.Default()

	//路由对象

	//绑定json
	r.POST("/loginJson", jsonHandler)

	//绑定form表单
	r.POST("/loginForm", formHandler)

	//绑定querystring
	r.GET("/loginForm", querystringHandle)

	//启动 端口8060
	r.Run(":8060")

}
