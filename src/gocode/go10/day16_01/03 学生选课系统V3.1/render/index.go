package render

import "github.com/gin-contrib/multitemplate"

func CreateMyRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	//登录页面
	//render.AddFromFiles("loginText.html", "templates/loginText.html")
	render.AddFromFiles("login.html", "templates/login.html")
	//首页
	render.AddFromFiles("index.html", "templates/base.html", "templates/index.html")
	render.AddFromFiles("student.html", "templates/base.html", "templates/student.html")
	//某一个学生的详情页
	render.AddFromFiles("detailOneStudent.html", "templates/base.html", "templates/detailOneStudent.html")
	//获取选课页面
	render.AddFromFiles("selectCourse.html", "templates/base.html", "templates/selectCourse.html")
	//添加学生
	render.AddFromFiles("getStuAdd", "templates/base.html", "templates/getStuAdd.html")
	//修改(更新)学生
	render.AddFromFiles("editStudent.html", "templates/base.html", "templates/editStudent.html")
	render.AddFromFiles("class", "templates/base.html", "templates/class.html")
	render.AddFromFiles("course", "templates/base.html", "templates/course.html")

	return render
}
