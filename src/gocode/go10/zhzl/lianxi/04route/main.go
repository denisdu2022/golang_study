package main

import "github.com/gin-gonic/gin"

func main() {
	//1.获取路由对象
	r := gin.Default()

	//3.加载模板文件
	r.LoadHTMLGlob("templates/*")

	//书籍路由
	//查看书籍
	r.GET("/books", func(ctx *gin.Context) {
		//响应
		ctx.JSON(200, gin.H{
			"msg": "查看书籍",
		})
	})

	//添加书籍
	r.POST("/books/add", func(ctx *gin.Context) {
		//相应
		ctx.JSON(200, gin.H{
			"msg": "添加书籍",
		})
	})

	//编辑书籍
	r.PUT("/books/edit", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "编辑书籍",
		})
	})

	//删除书籍
	r.DELETE("/books/delete", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "删除书籍",
		})
	})

	//NotRoute
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(200, "404.html", nil)
	})

	//2.启动
	r.Run()
}
