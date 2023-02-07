package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute(r *gin.Engine) {
	//书籍
	//InitBookRoute在同一个包下不需要引入,直接调用
	InitBookRoute(r)

	//出版社
	//InitPublishRoute在同一个包下不需要引入,直接调用
	InitPublishRoute(r)

	//NoRoute路由匹配不成功,404
	//没有路径,只有handler
	r.NoRoute(func(context *gin.Context) {
		//http.StatusNotFound就是404也可以直接写404
		context.HTML(http.StatusNotFound, "404.html", nil)
	})
}
