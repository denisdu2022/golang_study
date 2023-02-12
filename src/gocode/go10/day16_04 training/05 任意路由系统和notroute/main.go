package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//获取gin 路由对象
	r := gin.Default()

	//路由系统

	//任意路由
	r.Any("/book", func(ctx *gin.Context) {
		//获取请求方式
		method := ctx.Request.Method
		fmt.Println(method)
		//get 获取页面
		//post 登录

		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"请求方式是:": method,
		})
	})

	//NotRoute
	//路由匹配不成功返回404
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"404": "您访问的页面不存在",
		})
	})

	//启动,默认端口8080
	r.Run()
}
