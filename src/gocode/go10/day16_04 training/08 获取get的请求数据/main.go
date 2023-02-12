package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//获取路由对象
	r := gin.Default()

	//获取基本的请求信息

	//路由
	r.GET("/", func(ctx *gin.Context) {

		//1.请求
		//获取请求方式
		method := ctx.Request.Method
		fmt.Println(method)

		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"请求方式是:": method,
		})
	})

	//任意路由
	r.Any("/test", func(ctx *gin.Context) {
		////请求
		////2.获取请求路径
		//URL := ctx.Request.URL
		//urlPath := ctx.Request.URL.Path
		//fullPath := ctx.FullPath()
		//urlHost := ctx.Request.URL.Host
		//host := ctx.Request.Host
		//fmt.Println("URL:", URL)
		//fmt.Println("UrlPath:", urlPath)
		//fmt.Println("fullPath", fullPath)
		//fmt.Println("urlHost:", urlHost)
		//fmt.Println("host: ", host)
		//
		////响应
		//ctx.String(http.StatusOK, "test")

		//3.请求头
		//获取请求头
		//获取客户端信息 ua头(反爬虫可以通过ua头)
		ua := ctx.GetHeader("user-agent")
		//获取token
		token := ctx.GetHeader("Postman-Token")
		//获取缓存信息
		cache := ctx.GetHeader("Cache-Control")
		//判断是否缓存
		//if cache == "no-cache" {
		//	cache = "不缓存"
		//}
		//获取常见头
		//获取客户端IP
		cIP := ctx.ClientIP()
		//ip
		rIP := ctx.RemoteIP()
		//id+端口
		rAdd := ctx.Request.RemoteAddr

		//打印请求头
		fmt.Println("UA头: ", ua)
		fmt.Println("token: ", token)
		fmt.Println("cache: ", cache)
		fmt.Println("cIP: ", cIP)
		fmt.Println("rip: ", rIP)
		fmt.Println("rAdd: ", rAdd)

		//响应
		ctx.String(http.StatusOK, "test")
	})

	//启动
	r.Run()
}
