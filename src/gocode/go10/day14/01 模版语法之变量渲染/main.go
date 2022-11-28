package main

import "github.com/gin-gonic/gin"

func index(ctx *gin.Context) {
	//构建Go语言下的各类数据
	var name = "luobo"

	ctx.HTML(200, "index.html", gin.H{
		"name": name,
	})
}

func main() {
	//获取路由对象
	r := gin.Default()
	//加载模版文件
	r.LoadHTMLGlob("templates/*")
	//设置静态资源窗口
	r.Static("/static", "./static")

	r.GET("/", index)

	//启动
	r.Run()

}
