package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//学生管理路由函数

func getStu(ctx *gin.Context) {
	//响应返回学生管理
	//ctx.String(200, "学生管理")
	ctx.JSON(200, gin.H{
		"student": "学生管理",
	})
}

func main() {
	//获取路由对象
	r := gin.Default()

	//路由
	//1.站外重定向
	r.GET("/redirectBaidu", func(ctx *gin.Context) {
		//重定向到百度
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	//2.站内重定向
	r.GET("redirectStudent", func(ctx *gin.Context) {
		//重定向到学生管理
		ctx.Redirect(http.StatusMovedPermanently, "/student")
	})

	//学生管理路由
	r.GET("/student", getStu)

	//启动
	r.Run()
}
