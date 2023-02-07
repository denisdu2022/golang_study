package main

import "github.com/gin-gonic/gin"

func main() {
	//获取路由对象
	r := gin.Default()

	//响应html页面
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//路由
	r.GET("/test02", func(ctx *gin.Context) {
		//响应HTML页面
		ctx.HTML(200, "test02.html", nil)
	})

	//响应简单字符串
	r.GET("/test01", func(ctx *gin.Context) {
		//响应字符串
		ctx.String(200, "hello web")
	})

	//响应JSON数据
	r.GET("/test03", func(ctx *gin.Context) {
		//响应
		ctx.JSON(200, gin.H{
			"msg": "这是test03",
		})
	})

	//响应xml数据
	r.GET("/test04", func(ctx *gin.Context) {
		//响应xml
		ctx.XML(200, gin.H{
			"user_id":   "1001",
			"user_name": "test04",
		})
	})

	//启动
	r.Run()
}
