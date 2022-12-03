package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// 结构体
type Student struct {
	Sid    int
	Name   string
	Age    int
	Gender string
	Email  string
	Addr   string
}

func index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func student(ctx *gin.Context) {
	//模拟从数据库取出的数据
	//实例化结构体对象
	students := []Student{
		{1000, "张三", 22, "male", "111@qq.com", "河北省承德市"}, {1001, "李四", 21, "male", "222@qq.com", "河北省石家庄"}, {1002, "王五", 23, "male", "333@qq.com", "河北省廊坊"}, {1003, "赵六", 20, "male", "555@qq.com", "河北省承德市"},
	}
	ctx.HTML(200, "student.html", gin.H{
		"students": students,
	})
}

func class(ctx *gin.Context) {
	ctx.HTML(200, "class", nil)
}
func course(ctx *gin.Context) {
	ctx.HTML(200, "course", nil)
}

func createMyRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("index.html", "templates/base.html", "templates/index.html")
	render.AddFromFiles("student.html", "templates/base.html", "templates/student.html")
	render.AddFromFiles("class", "templates/base.html", "templates/class.html")
	render.AddFromFiles("course", "templates/base.html", "templates/course.html")

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
	//班级页面
	r.GET("/class", class)
	//课程管理
	r.GET("/course", course)

	//启动
	r.Run()

}
