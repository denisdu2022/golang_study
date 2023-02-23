package route

import (
	. "cssv2.3/core"
	"github.com/gin-gonic/gin"
)

func InitIndex(r *gin.Engine) {

	//首页

	r.GET("/", Index)

	r.GET("/index", GetIndex)

}
