package route

import "github.com/gin-gonic/gin"

func InitPublishRoute(r *gin.Engine) {

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
}
