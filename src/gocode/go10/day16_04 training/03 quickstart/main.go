package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//首页路由函数

func getIndex(ctx *gin.Context) {
	//响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"msg": "index",
	//})

	//响应页面
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func main() {

	//获取路由对象
	r := gin.Default()

	//加载模板文件
	r.LoadHTMLGlob("templates/*")

	//路由
	r.GET("/", getIndex)

	//启动 默认8080端口
	r.Run()
}
