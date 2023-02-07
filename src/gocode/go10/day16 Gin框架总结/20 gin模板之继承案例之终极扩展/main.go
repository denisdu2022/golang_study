package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

//定义学生的结构体

type Student struct {
	Sid    int
	Name   string
	Age    int
	Gender string
}

// 首页函数
func getIndex(ctx *gin.Context) {

	//响应
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// 学生管理函数
func getStudent(ctx *gin.Context) {
	//模拟数据,模拟在数据库中的数据
	//实例化学生对象并赋值
	student := []Student{
		{1001, "haha", 21, "男"},
		{1001, "liu", 22, "男"},
		{1001, "han", 20, "男"},
	}
	//响应
	ctx.HTML(http.StatusOK, "student.html", gin.H{
		"stu": student,
	})
}

// 课程管理
func getCourse(ctx *gin.Context) {
	//响应
	ctx.HTML(http.StatusOK, "course.html", nil)
}

// 班级管理
func getClass(ctx *gin.Context) {
	//响应
	ctx.HTML(http.StatusOK, "class.html", nil)
}

// 使用multitemplate
func createMyRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	//第一个参数与函数中的name相同,但是不可以写函数名   (母版在前,子版在后)第二个参数加载继承的模板.html   最后是具体的渲染页面
	//首页页面
	render.AddFromFiles("index.html", "templates/base.html", "templates/index.html")
	//学生管理页面
	render.AddFromFiles("student.html", "templates/base.html", "templates/student.html")
	//课程管理页面
	render.AddFromFiles("course.html", "templates/base.html", "templates/course.html")
	//班级管理
	render.AddFromFiles("class.html", "templates/base.html", "templates/class.html")
	return render
}

func main() {
	//获取路由对象
	r := gin.Default()
	//加载模板文件  gin框架默认使用的单模板
	//r.LoadHTMLGlob("templates/*")

	//使用multitemplate ,调用createMyRender函数
	r.HTMLRender = createMyRender()

	//开放静态文件窗口
	r.Static("/static", "./static")

	//路由
	//首页
	r.GET("/", getIndex)

	//学生页面
	r.GET("/student", getStudent)

	//课程页面
	r.GET("/course", getCourse)

	//班级页面
	r.GET("/class", getClass)

	//启动
	r.Run()
}
