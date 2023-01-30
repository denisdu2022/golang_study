package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//1.获取路由对象
	r := gin.Default()

	//获取基本的请求信息

	//get路由
	r.GET("/index", func(ctx *gin.Context) {
		//请求
		//1.基本请求信息
		method := ctx.Request.Method
		fmt.Println("请求方式: ", method)

		//响应
		ctx.JSON(200, gin.H{
			"请求方式:": method,
		})
	})

	//任意路由
	r.Any("/login", func(ctx *gin.Context) {
		//请求
		//获取基本的请求信息
		////1.请求方式
		//method := ctx.Request.Method
		////打印请求方式
		//fmt.Println("请求方式: ", method)
		//
		////响应
		//ctx.JSON(200, gin.H{
		//	"msg":        "请求成功!",
		//	"请求方式: ": method,
		//})

		////2.请求路径(URL)
		////请求
		//url := ctx.Request.URL
		//urlPath := ctx.Request.URL.Path
		//fullPath := ctx.FullPath()
		////打印请求路径
		//fmt.Println("URL: ", url)
		//fmt.Println("URLPath: ", urlPath)
		//fmt.Println("fullPath: ", fullPath)
		////响应
		//ctx.JSON(200, gin.H{
		//	"msg":      "请求成功!",
		//	"url":      url,
		//	"urlPath":  urlPath,
		//	"fullPath": fullPath,
		//})

		//3.请求头
		//获取客户端信息 UA头(反爬虫可以通过UA头)
		ua := ctx.GetHeader("user-agent")
		//获取token
		token := ctx.GetHeader("Postman-Token")
		//常见头
		//获取客户端IP
		cIP := ctx.ClientIP()
		//IP+端口
		rIP := ctx.RemoteIP()
		//ip
		rAdd := ctx.Request.RemoteAddr

		//打印请求头
		fmt.Println("ua头: ", ua)
		fmt.Println("token: ", token)
		fmt.Println("cIP: ", cIP)
		fmt.Println("rIP: ", rIP)
		fmt.Println("rAdd: ", rAdd)

		//响应
		ctx.JSON(200, gin.H{
			"ua头: ": ua,
			"token": token,
			"cIP":   cIP,
			"rIP":   rIP,
			"rAdd":  rAdd,
		})

	})

	//2.启动
	r.Run()

}
