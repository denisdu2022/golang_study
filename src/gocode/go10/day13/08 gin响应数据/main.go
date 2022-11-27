package main

import "github.com/gin-gonic/gin"

func main() {

	//Response 响应

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	//静态请求  静态请求是请求的静态资源
	//设置静态文件窗口  root 是物理路径  relativePath静态路由,不在进入下边的动态路由中
	r.Static("/static", "./static")

	//以下是动态请求
	r.GET("/ResponseString", func(context *gin.Context) {
		context.String(200, "ResponseString!")
	})

	r.GET("/ResponseHtml", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})

	r.GET("/ResponseJson", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"user_id":   1001,
			"user_name": "luobo",
		})
	})

	r.GET("/ResponseXml", func(context *gin.Context) {
		context.XML(200, gin.H{
			"user_id":   1001,
			"user_name": "luobo",
			"friends":   []string{"a", "b", "c"},
		})
	})

	r.GET("/login", func(context *gin.Context) {

		context.HTML(200, "login.html", nil)
	})

	r.GET("/Bootstrap", func(context *gin.Context) {
		context.HTML(200, "Bootstrap.html", nil)
	})

	r.Run(":8060")
}
