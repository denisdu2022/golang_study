package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//获取路由对象
	r := gin.Default()

	//获取基本的请求信息

	//r.GET("/index", func(context *gin.Context) {
	//	//请求
	//	//1.基本请求信息
	//	//打印请求方式
	//	fmt.Println("请求方式: ", context.Request.Method)
	//
	//	//响应
	//	context.JSON(200, gin.H{
	//		"msg":  "请求成功",
	//		"请求方式": context.Request.Method,
	//	})
	//})

	r.Any("/index", func(context *gin.Context) {
		//请求
		//1.基本请求信息
		////打印请求方式
		//fmt.Println("请求方式: ", context.Request.Method)
		//
		////响应
		//context.JSON(200, gin.H{
		//	"msg":  "请求成功",
		//	"请求方式": context.Request.Method,
		//})

		////打印请求路径(URL)
		//fmt.Println("请求URL: ", context.Request.URL)
		////或者
		//fmt.Println("请求URL: ", context.Request.URL.Path)
		////或者
		//fmt.Println("请求URL: ", context.FullPath())

		////响应
		//context.JSON(200, gin.H{
		//	"msg":    "请求成功",
		//	"请求URL":  context.Request.URL,
		//	"请求URL1": context.Request.URL.Path,
		//	"请求URL2": context.FullPath(),
		//})

		//2.请求头
		//获取客户端信息(反爬虫可以通过判断UA头)
		fmt.Println(context.GetHeader("user-agent"))
		//获取客户端Token
		fmt.Println(context.GetHeader("Postman-Token"))
		//常见头
		//获取客户端IP
		fmt.Println("ClientIP:", context.ClientIP())
		//IP不带端口
		fmt.Println("RemoteIP", context.RemoteIP())
		//IP带端口
		fmt.Println("RemoteIP+Port:", context.Request.RemoteAddr)

		//响应
		context.JSON(200, gin.H{
			"客户端信息":         context.GetHeader("user-agent"),
			"获取客户端Token":    context.GetHeader("Postman-Token"),
			"ClientIP":      context.ClientIP(),
			"RemoteIP":      context.RemoteIP(),
			"RemoteIP+Port": context.Request.RemoteAddr,
		})
	})

	//启动
	r.Run()
}
