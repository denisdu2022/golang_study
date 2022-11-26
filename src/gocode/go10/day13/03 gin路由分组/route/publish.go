package route

import "github.com/gin-gonic/gin"

func InitPublishRoute(r *gin.Engine) {
	//publish
	publishRouter := r.Group("/publish")

	//Get请求
	publishRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "查看所有出版社",
		})
	})

	//POST请求
	publishRouter.POST("/add", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "添加出版社",
		})
	})

	publishRouter.GET("/edit", func(c *gin.Context) {
		//删除书籍功能

		c.JSON(200, gin.H{
			"message": "编辑出版社",
		})
	})

	publishRouter.GET("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "删除出版社",
		})
	})
}
