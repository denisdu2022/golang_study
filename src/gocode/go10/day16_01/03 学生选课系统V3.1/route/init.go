package route

import (
	. "css/core"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	//登录
	r.GET("/login", GetLoginHtml)
	r.POST("/login", Login)
	//首页
	r.GET("/", Index)
	//学生路由组
	student := r.Group("/student")
	InitStudentRoute(student)
	//班级页面
	r.GET("/class", GetClass)
	//课程管理
	r.GET("/course", GetCourse)
}
