package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func student(ctx *gin.Context) {
	ctx.HTML(200, "student.html", nil)
}

func createMyRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("index.html", "templates/base.html", "templates/index.html")
	render.AddFromFiles("student.html", "templates/base.html", "templates/student.html")

	return render
}

func main() {
	//获取路由对象
	r := gin.Default()
	////加载模版文件
	//r.LoadHTMLGlob("templates/*")

	//加载模版文件
	r.HTMLRender = createMyRender()

	//设置静态文件资源窗口
	r.Static("/static", "./static")

	//首页
	r.GET("/", index)
	//学生页面
	r.GET("/student", student)

	//启动
	r.Run()

}
