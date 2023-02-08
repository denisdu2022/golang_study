package route

import (
	. "css/core"
	"github.com/gin-gonic/gin"
)

func InitStudentRoute(r *gin.RouterGroup) {

	//学生首页页面
	r.GET("/", GetAllStudent)
	//学生个人详情页
	r.GET("/:sno/", GetOneStudent)
	//学生个人选课系统
	r.GET("/:sno/selectCourse/", GetSelectCourse)
	r.POST("/:sno/selectCourse/", SelectCourse)
	//添加学生
	r.GET("/add", GetStuAdd)
	r.POST("/add", PostStuAdd)
	//删除学生
	r.GET("/delete/:delID", DeleteStu)
	//编辑学生
	r.GET("/edit/:editID", GetStuEditHtml)
}
