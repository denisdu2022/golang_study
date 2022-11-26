package route

import (
	. "ginRouterZu/core"
	"github.com/gin-gonic/gin"
)

func InitBookRoute(r *gin.Engine) {

	//book
	bookRouter := r.Group("/book")

	//Get请求
	bookRouter.GET("/", Book)

	//POST请求
	bookRouter.POST("/add", AddBook)

	bookRouter.GET("/edit", EditBook)

	bookRouter.GET("/delete", DeleteBook)
}
