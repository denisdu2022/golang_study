package route

import (
	. "cssv2.3/core"
	"github.com/gin-gonic/gin"
)

func InitStuRoute(r *gin.Engine) {

	//学生路由组
	stuRoute := r.Group("/student")

	//查看学生
	stuRoute.GET("/", GetStudent)

	//查看学生的个人详细信息
	stuRoute.GET("/:sno", GetOneStuHTml)

	//学生选课页面
	stuRoute.GET("/selectCourse/:sno", GetSelectCourseHtml)

	//学生选课提交
	stuRoute.POST("/selectCourse/:sno", PostSelectCourse)

	//获取添加学生页面
	stuRoute.GET("/add", GetStudentAddHtml)

	//添加学生
	stuRoute.POST("/add", PostStudentAdd)

	//删除学生
	stuRoute.GET("/delete/:delID", GetDeleteStudent)

	//编辑学生
	stuRoute.GET("/edit/:editID", GetEditStudent)
}
