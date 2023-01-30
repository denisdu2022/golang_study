package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute(r *gin.Engine) {

	//路由分组

	//调用book路由
	InitBookRoute(r)

	//调用publish路由
	InitPublishRoute(r)

	//NotRoute 404
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", nil)
	})

}
