package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitBookRoute(r *gin.Engine) {
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

}
