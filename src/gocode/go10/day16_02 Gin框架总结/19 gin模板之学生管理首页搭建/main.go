package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 首页函数
func getIndex(ctx *gin.Context) {
	//响应
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	//获取路由对象
	r := gin.Default()
	//加载模板文件  gin框架默认使用的单模板
	r.LoadHTMLGlob("templates/*")
	//开放静态文件窗口
	r.Static("/static", "./static")

	//路由
	//首页
	r.GET("/", getIndex)

	//启动
	r.Run()
}
