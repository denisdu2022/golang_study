package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//获取路由对象
	r := gin.Default()

	//加载模板文件
	r.LoadHTMLGlob("templates/*")

	//Get查看书籍
	r.GET("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "查看书籍!",
		})
	})

	//Post添加书籍
	r.POST("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "添加数据!",
		})
	})

	//Put修改书籍
	r.PUT("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "更新书籍!",
		})
	})

	//Delete删除书籍
	r.DELETE("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "删除书籍!",
		})
	})

	//NoRoute路由匹配不成功,404
	r.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	//启动
	r.Run(":8090")
}
