package route

import (
	. "cssv2.3/core"
	"github.com/gin-gonic/gin"
)

func InitLogin(r *gin.Engine) {
	//登录

	r.GET("/login", GetLoginHtml)

	r.POST("/login", PostLogin)

}
