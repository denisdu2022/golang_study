package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute(r *gin.Engine) {

	//路由分组

	//书籍路由
	InitBookRoute(r)

	//出版社路由
	InitPublishRoute(r)

	//NotRoute
	//路由匹配不成功返回404
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"404": "您访问的页面不存在",
		})
	})
}
