package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

//学生结构体

type Student struct {
	Sid    int
	Name   string
	Age    int
	Gender string
}

// 首页路由函数
func getIndex(ctx *gin.Context) {
	//响应首页页面
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// 学生路由函数
func getStudent(ctx *gin.Context) {
	//模拟数据库取出的数据
	//实例化学生对象
	students := []Student{
		{1001, "张三", 20, "男"},
		{1002, "李四", 21, "女"},
		{1003, "王五", 22, "男"},
	}
	//响应学生页面
	ctx.HTML(http.StatusOK, "student.html", gin.H{
		"students": students,
	})
}

// 课程路由函数
func getCourse(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "course.html", nil)
}

// 班级路由函数
func getClass(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "class.html", nil)
}

// multitemplate函数
func createMyRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	//首页页面
	render.AddFromFiles("index.html", "templates/base.html", "templates/index.html")
	//学生页面
	render.AddFromFiles("student.html", "templates/base.html", "templates/student.html")
	//课程页面
	render.AddFromFiles("course.html", "templates/base.html", "templates/course.html")
	//班级页面
	render.AddFromFiles("class.html", "templates/base.html", "templates/class.html")
	return render
}

func main() {

	//获取路由对象
	r := gin.Default()

	//使用HTMLRender调用createMyRender函数
	r.HTMLRender = createMyRender()

	//开放静态文件窗口
	r.Static("/static", "./static")

	//路由
	r.GET("/", getIndex)
	//学生管理
	r.GET("/student", getStudent)
	//课程路由
	r.GET("/course", getCourse)
	//班级路由
	r.GET("/class", getClass)

	//启动
	r.Run()

}
