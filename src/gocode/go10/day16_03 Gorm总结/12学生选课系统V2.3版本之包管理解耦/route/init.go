package route

import (
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	//路由

	//首页路由
	InitIndex(r)

	//登录路由
	InitLogin(r)

	//学生路由
	InitStuRoute(r)

	//课程路由
	InitCourseRoute(r)

	//班级路由
	InitClassRoute(r)
}
