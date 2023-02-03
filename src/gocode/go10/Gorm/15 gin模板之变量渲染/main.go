package main

import "github.com/gin-gonic/gin"

func main() {
	//获取路由对象
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//开放静态文件窗口
	r.Static("/static", "./static")

	//启动
	r.Run()
}
