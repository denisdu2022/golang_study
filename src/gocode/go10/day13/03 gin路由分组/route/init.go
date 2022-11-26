package route

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine) {

	//book
	InitBookRoute(r)

	//publish
	InitPublishRoute(r)

	//加载模版文件
	r.LoadHTMLGlob("templates/*")

	//NO router
	r.NoRoute(func(context *gin.Context) {
		context.HTML(404, "404.html", nil)
	})
}
