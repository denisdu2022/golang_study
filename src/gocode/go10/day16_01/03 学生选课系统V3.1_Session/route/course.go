package route

import (
	. "css/core"
	"github.com/gin-gonic/gin"
)

func InitCourseRoute(r *gin.RouterGroup) {
	//课程首页页面
	r.GET("/", GetCourse)
}
