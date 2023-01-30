package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	//获取路由对象
	r := gin.Default()

	//加载模板文件
	r.LoadHTMLGlob("templates/*")

	//login路由
	//get拿页面
	r.GET("/login", func(ctx *gin.Context) {
		//响应
		ctx.HTML(200, "login.html", nil)
	})

	//post登录
	r.POST("/login", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "登录成功",
		})
	})

	//任意路由
	r.Any("/books", func(ctx *gin.Context) {
		//获取请求路径
		path := ctx.FullPath()
		fmt.Println(path)

		//获取请求方式
		req := ctx.Request
		fmt.Println(req)

		//响应
		ctx.JSON(200, gin.H{
			"请求路径": path,
		})

		/*
			请求方式
			&{POST /books HTTP/1.1 1 1..
			&{GET /books HTTP/1.1 1 1..
		*/
	})

	//NotRoute 路由匹配不成功返回404
	//NotRoute没有路径只有handler
	r.NoRoute(func(ctx *gin.Context) {
		ctx.String(200, "Not Fount 404!!!")
	})

	//启动
	r.Run(":8090")

}
