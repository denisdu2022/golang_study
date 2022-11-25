package main

import "github.com/gin-gonic/gin"

func main() {

	//获取路由对象
	r := gin.Default()

	//路由分组
	bookRouter := r.Group("/books")

	//Get请求
	bookRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "查看所有书籍",
		})
	})

	//POST请求
	bookRouter.POST("/add", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "添加书籍",
		})
	})

	bookRouter.GET("/edit", func(c *gin.Context) {
		//删除书籍功能

		c.JSON(200, gin.H{
			"message": "编辑书籍",
		})
	})

	bookRouter.GET("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "删除书籍",
		})
	})

	//加载模版文件
	r.LoadHTMLGlob("templates/*")

	//NO router
	r.NoRoute(func(context *gin.Context) {
		context.HTML(404, "404.html", nil)
	})

	//启动
	r.Run()
}
