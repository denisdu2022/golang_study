package route

import "github.com/gin-gonic/gin"

func InitBookRoute(r *gin.Engine) {

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
			"msg": "添加书籍!",
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
}
