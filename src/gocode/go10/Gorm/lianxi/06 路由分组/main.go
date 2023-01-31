package main

import "github.com/gin-gonic/gin"

func main() {
	//获取路由对象
	r := gin.Default()

	////书籍路由
	////查看书籍
	//r.GET("/books", func(ctx *gin.Context) {
	//	//响应
	//	ctx.JSON(200, gin.H{
	//		"msg": "查看书籍",
	//	})
	//})
	//
	////添加书籍
	//r.POST("/books/add", func(ctx *gin.Context) {
	//	//相应
	//	ctx.JSON(200, gin.H{
	//		"msg": "添加书籍",
	//	})
	//})
	//
	////编辑书籍
	//r.PUT("/books/edit", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{
	//		"msg": "编辑书籍",
	//	})
	//})
	//
	////删除书籍
	//r.DELETE("/books/delete", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{
	//		"msg": "删除书籍",
	//	})
	//})
	//
	//
	////出版社路由
	////查看出版社
	//r.GET("/publishes", func(ctx *gin.Context) {
	//	//响应
	//	ctx.JSON(200, gin.H{
	//		"msg": "查看出版社",
	//	})
	//})
	//
	////添加出版社
	//r.POST("/publishes/add", func(ctx *gin.Context) {
	//	//相应
	//	ctx.JSON(200, gin.H{
	//		"msg": "添加出版社",
	//	})
	//})
	//
	////编辑出版社
	//r.PUT("/publishes/edit", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{
	//		"msg": "编辑出版社",
	//	})
	//})
	//
	////删除出版社
	//r.DELETE("/publishes/delete", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{
	//		"msg": "删除出版社",
	//	})
	//})

	//路由分组
	//获取书籍路由组对象
	booksRoute := r.Group("/books")

	//书籍路由
	//查看书籍
	booksRoute.GET("/", func(ctx *gin.Context) {
		//响应
		ctx.JSON(200, gin.H{
			"msg": "查看书籍",
		})
	})

	//添加书籍
	booksRoute.POST("/add", func(ctx *gin.Context) {
		//相应
		ctx.JSON(200, gin.H{
			"msg": "添加书籍",
		})
	})

	//编辑书籍
	booksRoute.PUT("/edit", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "编辑书籍",
		})
	})

	//删除书籍
	booksRoute.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "删除书籍",
		})
	})

	//出版社路由
	//获取出版社路由对象
	publishes := r.Group("/publishes")
	//查看出版社
	publishes.GET("/", func(ctx *gin.Context) {
		//响应
		ctx.JSON(200, gin.H{
			"msg": "查看出版社",
		})
	})

	//添加出版社
	publishes.POST("/add", func(ctx *gin.Context) {
		//相应
		ctx.JSON(200, gin.H{
			"msg": "添加出版社",
		})
	})

	//编辑出版社
	publishes.PUT("/edit", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "编辑出版社",
		})
	})

	//删除出版社
	publishes.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "删除出版社",
		})
	})

	//启动
	r.Run()
}
