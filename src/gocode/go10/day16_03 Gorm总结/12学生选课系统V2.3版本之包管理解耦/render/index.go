package render

import "github.com/gin-contrib/multitemplate"

// multitemplate函数

func CreateMyRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	//登录页面
	render.AddFromFiles("login.html", "templates/login.html")
	//首页页面
	render.AddFromFiles("index.html", "templates/base.html", "templates/index.html")
	//学生页面
	render.AddFromFiles("student.html", "templates/base.html", "templates/student.html")
	//学生个人详情
	render.AddFromFiles("detailStudent.html", "templates/base.html", "templates/detailStudent.html")
	//学生选课页面
	render.AddFromFiles("selectCourse.html", "templates/base.html", "templates/selectCourse.html")
	//添加学生
	render.AddFromFiles("getStuAdd.html", "templates/base.html", "templates/getStuAdd.html")
	//编辑学生
	render.AddFromFiles("getEditStu.html", "templates/base.html", "templates/getEditStu.html")
	//课程页面
	render.AddFromFiles("course.html", "templates/base.html", "templates/course.html")
	//添加课程
	render.AddFromFiles("getCourseAdd.html", "templates/base.html", "templates/getCourseAdd.html")
	//编辑课程
	render.AddFromFiles("editCourse.html", "templates/base.html", "templates/editCourse.html")
	//班级页面
	render.AddFromFiles("class.html", "templates/base.html", "templates/class.html")
	return render
}
