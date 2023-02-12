package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitPublishRoute(r *gin.Engine) {
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
}
