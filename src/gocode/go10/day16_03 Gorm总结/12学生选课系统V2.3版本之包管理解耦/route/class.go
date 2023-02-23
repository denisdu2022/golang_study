package route

import (
	. "cssv2.3/core"
	"github.com/gin-gonic/gin"
)

func InitClassRoute(r *gin.Engine) {
	//班级路由
	r.GET("/class", GetClass)
}
