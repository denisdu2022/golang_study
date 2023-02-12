package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//获取gin 路由对象
	r := gin.Default()

	//路由系统

	//路由分组
	//书籍路由
	bookRoute := r.Group("/books")
	//get
	bookRoute.GET("/", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "查看书籍",
		})
	})
	//post
	bookRoute.POST("/", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "添加书籍",
		})
	})
	//put
	bookRoute.PUT("/edit", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "修改书籍",
		})
	})
	//delete
	bookRoute.DELETE("/delete", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "删除书籍",
		})
	})

	//出版社路由
	publishes := r.Group("/publish")
	//get
	publishes.GET("/", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "查看出版社",
		})
	})
	//post
	publishes.POST("/", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "添加出版社",
		})
	})
	//修改出版社
	publishes.GET("/edit", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "修改出版社",
		})
	})
	//删除出版社
	publishes.GET("/delete", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "删除出版社",
		})
	})

	//NotRoute
	//路由匹配不成功返回404
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"404": "您访问的页面不存在",
		})
	})

	//启动,默认端口8080
	r.Run()
}
