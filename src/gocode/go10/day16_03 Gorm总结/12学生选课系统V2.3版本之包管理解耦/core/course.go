package core

import (
	. "cssv2.3/database"
	. "cssv2.3/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 课程路由函数

func GetCourse(ctx *gin.Context) {
	//获取课程名
	coName := ctx.Query("coName")
	//从数据库查询所有课程
	var courses []Course

	//基于模糊查询的搜索框
	if coName == "" {
		//查询
		DB.Find(&courses)

		//响应 查出的数据传给前端
		ctx.HTML(http.StatusOK, "course.html", gin.H{
			"courses": courses,
		})
	} else {
		//查询
		DB.Where("name like ?", "%"+coName+"%").Find(&courses)
		fmt.Println(courses)

		//响应 查出的数据传给前端
		ctx.HTML(http.StatusOK, "course.html", gin.H{
			"courses": courses,
		})
	}

}

//添加课程
//GET拿页面

func GetCourseAddHtml(ctx *gin.Context) {
	//查询教师信息
	var teachers []Teacher
	//数据库查询
	DB.Find(&teachers)
	//响应添加课程页面
	ctx.HTML(http.StatusOK, "getCourseAdd.html", gin.H{
		"teachers": teachers,
	})
}

//post提交

func PostCourseAdd(ctx *gin.Context) {
	//获取前端传来的参数
	//课程名
	coName := ctx.PostForm("coName")
	//学分  获取到的是string类型,需要类型转换
	strCredit := ctx.PostForm("credit")
	intCredit, _ := strconv.Atoi(strCredit)
	//周期
	strPeriod := ctx.PostForm("period")
	intPeriod, _ := strconv.Atoi(strPeriod)
	//任课老师ID
	strTID := ctx.PostForm("tID")
	intTID, _ := strconv.Atoi(strTID)

	//实例化课程对象并赋值
	course := Course{BaseModel: BaseModel{Name: coName}, Credit: intCredit, Period: intPeriod, TeacherID: intTID}
	//添加到数据库
	DB.Create(&course)

	//响应,重定向到课程管理页面
	ctx.Redirect(http.StatusMovedPermanently, "/course")

}

//删除课程

func GetDeleteCourse(ctx *gin.Context) {
	//获取delID
	delID := ctx.Param("delID")
	//先查询,在删除
	DB.Where("ID = ?", delID).Delete(&Course{})

	//响应
	//重定向到课程管理页面
	ctx.Redirect(http.StatusMovedPermanently, "/course")
}

//编辑课程

func GetEditCourseHtml(ctx *gin.Context) {
	//获取editID
	editID := ctx.Param("editID")
	//数据库查询
	var course Course
	DB.Where("ID = ?", editID).Find(&course)

	//查询teacher表
	var teachers []Teacher
	DB.Find(&teachers)

	//响应编辑页面
	ctx.HTML(http.StatusOK, "editCourse.html", gin.H{
		"course":   course,
		"teachers": teachers,
	})

}
