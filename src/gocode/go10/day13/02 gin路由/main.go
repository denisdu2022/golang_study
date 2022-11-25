package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//获取路由对象
	r := gin.Default()

	//Get请求
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "查看所有书籍",
		})
	})

	//POST请求
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "添加书籍",
		})
	})

	//PUT请求
	r.PUT("/book", func(c *gin.Context) {
		//删除书籍功能

		c.JSON(200, gin.H{
			"message": "编辑书籍",
		})
	})

	//DELTE请求
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "删除书籍",
		})
	})

	//r.GET("/login", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"msg": "get login",
	//	})
	//})
	//
	//r.POST("/login", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"msg": "post login",
	//	})
	//})

	r.Any("/login", func(context *gin.Context) {
		fmt.Println("path", context.FullPath())
		context.JSON(200, gin.H{
			"msg": "any login",
		})
	})

	//加载模版文件
	r.LoadHTMLGlob("templates/*")

	//NO router
	r.NoRoute(func(context *gin.Context) {
		context.HTML(404, "404.html", nil)
	})

	//启动
	r.Run(":8060")
}
