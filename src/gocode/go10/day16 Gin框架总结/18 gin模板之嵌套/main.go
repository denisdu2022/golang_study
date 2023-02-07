package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// test01
func getTest01(ctx *gin.Context) {
	//响应
	ctx.HTML(http.StatusOK, "t1.html", nil)
}

// test02
func getTest02(ctx *gin.Context) {
	//响应
	ctx.HTML(http.StatusOK, "t2.html", nil)
}

func main() {
	//获取路由对象
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//开放静态文件窗口
	r.Static("/static", "./static")

	//路由
	//test01
	r.GET("/test01", getTest01)
	//test02
	r.GET("/test02", getTest02)

	//启动
	r.Run()
}
