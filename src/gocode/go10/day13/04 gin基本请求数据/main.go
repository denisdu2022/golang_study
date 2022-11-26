package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//请求数据

	r := gin.Default()

	//获取基本的请求信息
	r.Any("/index", func(context *gin.Context) {
		//请求

		//1.基本请求信息
		//请求方式
		//fmt.Println("请求方式 Method: ", context.Request.Method)
		//req := context.Request.Method
		//请求的URL
		fmt.Println("请求路径 URL: ", context.Request.URL)
		fmt.Println("请求路径 FullPath: ", context.FullPath())
		url := context.Request.URL
		fullPath := context.FullPath()
		//请求头
		//user-agent
		fmt.Println("请求头 user-agent: ", context.GetHeader("user-agent"))
		userAgent := context.GetHeader("user-agent")
		//content-type
		fmt.Println("请求头 content-type: ", context.GetHeader("content-type"))
		contentType := context.GetHeader("content-type")
		//Postman-Token
		fmt.Println("请求头 Postman-Token: ", context.GetHeader("Postman-Token"))
		postmanToken := context.GetHeader("Postman-Token")
		//常见的请求头
		//远程IP和端口
		fmt.Println("RemoteIP: ", context.Request.RemoteAddr)
		remoteIp := context.Request.RemoteAddr
		fmt.Println("Host: ", context.GetHeader("Host"))
		host := context.GetHeader("Host")
		//客户端IP
		fmt.Println("ClientIP: ", context.ClientIP())
		clientIp := context.ClientIP()
		//测试
		fmt.Println("所有请求头: ", context.Request.Header)
		//响应
		//请求方式响应
		//context.JSON(200, gin.H{
		//	"REQ:": req,
		//})
		//请求路径响应
		context.JSON(200, gin.H{
			"url:":          url,
			"FullPath":      fullPath,
			"userAgent":     userAgent,
			"contentType":   contentType,
			"Postman-Token": postmanToken,
			"RemoteIp":      remoteIp,
			"Host":          host,
			"ClientIp":      clientIp,
		})
	})

	//启动
	r.Run(":8060")
}
