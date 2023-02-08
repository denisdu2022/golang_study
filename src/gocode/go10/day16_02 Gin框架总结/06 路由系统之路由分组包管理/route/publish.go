package route

import "github.com/gin-gonic/gin"

func InitPublishRoute(r *gin.Engine) {
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
}
