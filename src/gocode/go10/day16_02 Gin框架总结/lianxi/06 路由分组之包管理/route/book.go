package route

import "github.com/gin-gonic/gin"

func InitBookRoute(r *gin.Engine) {
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
}
