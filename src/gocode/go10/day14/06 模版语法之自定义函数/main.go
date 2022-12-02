package main

import (
	"github.com/gin-gonic/gin"
)

func test01(ctx *gin.Context) {
	ctx.HTML(200, "t1.html", nil)
}

func test02(ctx *gin.Context) {
	ctx.HTML(200, "t2.html", nil)
}

func main() {
	//获取路由对象
	r := gin.Default()
	//加载模版文件
	r.LoadHTMLGlob("templates/*")
	//设置静态文件资源窗口
	r.Static("/static", "./static")

	r.GET("/test01", test01)
	r.GET("/test02", test02)

	//启动
	r.Run()

}
