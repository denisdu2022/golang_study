package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//获取gin 路由对象
	r := gin.Default()

	//路由系统
	//书籍路由
	//get请求    使用匿名函数
	r.GET("/book", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "查看书籍",
		})
	})

	//post
	r.POST("/book", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "添加书籍",
		})
	})

	//put
	r.PUT("/book", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "修改书籍",
		})
	})

	//delete
	r.DELETE("/book", func(ctx *gin.Context) {
		//响应
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "删除书籍",
		})
	})

	//启动,默认端口8080
	r.Run()
}
