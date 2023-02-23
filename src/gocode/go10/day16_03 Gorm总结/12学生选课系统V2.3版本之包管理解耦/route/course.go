package route

import (
	. "cssv2.3/core"
	"github.com/gin-gonic/gin"
)

func InitCourseRoute(r *gin.Engine) {
	//课程路由组
	courseRoute := r.Group("/course")

	//查看课程
	courseRoute.GET("/", GetCourse)

	//添加课程页面 Get 拿页面  Post提交课程
	courseRoute.GET("/add", GetCourseAddHtml)

	//添加课程
	courseRoute.POST("/add", PostCourseAdd)

	//删除课程
	courseRoute.GET("/delete/:delID", GetDeleteCourse)

	//编辑课程
	courseRoute.GET("/edit/:editID", GetEditCourseHtml)

	//提交课程
}
