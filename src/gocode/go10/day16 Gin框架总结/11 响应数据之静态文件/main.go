package main

import (
	"github.com/gin-gonic/gin"
)

// 注册登录首页函数
func getIndex(ctx *gin.Context) {
	//响应页面
	ctx.HTML(200, "index.html", nil)
}

func main() {
	//获取路由对象
	r := gin.Default()

	//加载模板文件
	r.LoadHTMLGlob("templates/*")

	//静态文件配置
	//静态请求, 静态请求是请求的静态资源
	//设置静态文件窗口,relativePath是静态路由,不在进入下边的动态路由中,root是物理路径,文件夹所在位置
	r.Static("/static", "./static")

	//路由
	//注册登录首页
	r.GET("/", getIndex)

	//启动
	r.Run()
}
