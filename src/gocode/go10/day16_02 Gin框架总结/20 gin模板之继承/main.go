package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 首页函数
func getIndex(ctx *gin.Context) {
	//响应
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// 学生管理函数
func getStudent(ctx *gin.Context) {

	//响应
	ctx.HTML(http.StatusOK, "student", nil)
}

// 课程管理
func getCourse(ctx *gin.Context) {
	//响应
	ctx.HTML(http.StatusOK, "course", nil)
}

// 使用multitemplate
func createMyRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	//第一个参数与函数中的name相同,但是不可以写函数名   第二个参数加载继承的模板.html   最后是具体的渲染页面
	render.AddFromFiles("index.html", "templates/base.html", "templates/index.html")
	render.AddFromFiles("student", "templates/base.html", "templates/student.html")
	render.AddFromFiles("course", "templates/base.html", "templates/course.html")
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

	//启动
	r.Run()
}
