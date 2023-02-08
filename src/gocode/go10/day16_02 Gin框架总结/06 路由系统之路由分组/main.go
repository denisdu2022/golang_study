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

	////书籍相关
	////Get查看书籍
	//r.GET("/book", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"msg": "查看书籍!",
	//	})
	//})
	//
	////Post添加书籍
	//r.POST("/book/add", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"msg": "添加数据!",
	//	})
	//})
	//
	////修改书籍
	//r.GET("/book/edit", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"msg": "更新书籍!",
	//	})
	//})
	//
	////删除书籍
	//r.GET("/book/delete", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"msg": "删除书籍!",
	//	})
	//})

	//路由分组
	//书籍相关
	//r.Group获取路由组对象
	bookRoute := r.Group("/book")

	//书籍相关
	//Get查看书籍
	bookRoute.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "查看书籍!",
		})
	})

	//Post添加书籍
	bookRoute.POST("/add", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "添加数据!",
		})
	})

	//修改书籍
	bookRoute.GET("/edit", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "更新书籍!",
		})
	})

	//删除书籍
	bookRoute.GET("/delete", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "删除书籍!",
		})
	})

	//出版社相关
	//r.Group获取路由组对象
	publishRoute := r.Group("/publish")

	//出版社相关
	//Get查看出版社
	publishRoute.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "查看出版社!",
		})
	})

	//Post添加出版社
	publishRoute.POST("/add", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "添加出版社!",
		})
	})

	//修改出版社
	publishRoute.GET("/edit", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "修改出版社!",
		})
	})

	//删除出版社
	publishRoute.GET("/delete", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "删除出版社!",
		})
	})

	//NoRoute路由匹配不成功,404
	//没有路径,只有handler
	r.NoRoute(func(context *gin.Context) {
		//http.StatusNotFound就是404也可以直接写404
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	//启动
	r.Run(":8090")
}
