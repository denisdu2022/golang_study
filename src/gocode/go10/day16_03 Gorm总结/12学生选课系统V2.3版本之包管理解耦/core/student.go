package core

import (
	. "cssv2.3/database"
	. "cssv2.3/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// 学生路由函数

func GetStudent(ctx *gin.Context) {
	//基于模糊查询的搜索框实现
	//获取get请求参数
	name := ctx.Query("name")
	//打印获取到的参数
	fmt.Println("name>>> ", name) //name>>>  顾

	//定义student
	var student []Student

	//数据库查询
	//如果没有输入搜索的参数则向前端展示所有的数据
	if name == "" {
		//查询数据库中student表中的数据
		//查询所有学生
		//从数据库查询并写入student
		//连续跨表预加载两个表
		DB.Preload("Class.Tutor").Find(&student)
		//响应学生页面
		ctx.HTML(http.StatusOK, "student.html", gin.H{
			"student": student,
		})
	} else {
		//数据库查询
		//模糊查询 "name like ?", "%"+name+"%"  包含name参数的
		DB.Where("name like ?", "%"+name+"%").Preload("Class.Tutor").Find(&student)
		//响应学生页面
		ctx.HTML(http.StatusOK, "student.html", gin.H{
			"student": student,
		})
	}

}

//查看学生信息详情

func GetOneStuHTml(ctx *gin.Context) {
	//获取动态参数
	strSno := ctx.Param("sno")
	//强转
	intSno, _ := strconv.Atoi(strSno)

	//定义学生信息
	var stu Student
	//数据库查询 预加载班级表
	DB.Where("sno = ? ", intSno).Preload("Class").Take(&stu)

	//查询学生的关联课程
	var selectCourses []Course
	DB.Preload("Teacher").Model(&stu).Association("Courses").Find(&selectCourses)

	////响应
	//ctx.JSON(200, gin.H{
	//	"intSno": stu,
	//})

	//响应学生个人详情页面
	ctx.HTML(http.StatusOK, "detailStudent.html", gin.H{
		"stu":           stu,
		"selectCourses": selectCourses,
	})

	//fmt.Println("stu >>> ", stu)

}

// 学生选课
//获取选课页面

func GetSelectCourseHtml(ctx *gin.Context) {
	//获取学生学号
	strSno := ctx.Param("sno")
	intSno, _ := strconv.Atoi(strSno)

	//数据库查询
	var stu Student
	DB.Where("sno = ?", intSno).Preload("Class").Take(&stu)

	//获取所有课程
	var courses []Course
	DB.Preload("Teacher").Find(&courses)

	//响应选课页面
	ctx.HTML(http.StatusOK, "selectCourse.html", gin.H{
		"stu":     stu,
		"courses": courses,
	})
}

//提交选课

func PostSelectCourse(ctx *gin.Context) {
	//获取学生学号
	sno := ctx.Param("sno")
	fmt.Println("sno>>> ", sno)

	//数据库查询学生对象
	var stu Student
	DB.Where("sno = ?", sno).Take(&stu)

	//获取前端传来的数据 获取多个值,键值对使用PostFormArray
	courseIDSlice := ctx.PostFormArray("course_id")
	fmt.Println(courseIDSlice)

	//获取所有课程
	var courses []Course
	DB.Where("id in ?", courseIDSlice).Find(&courses)
	fmt.Println(courses)

	//绑定课程
	DB.Model(&stu).Association("Courses").Append(courses)

	//响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"courseID": courseIDSlice,
	//	"courses":  courses,
	//})
	//重定向跳转到个人详情页
	ctx.Redirect(http.StatusMovedPermanently, "/student/"+sno)
}

//添加学生路由函数

func GetStudentAddHtml(ctx *gin.Context) {
	//查询班级信息
	var classes []Class
	//数据库查询
	DB.Find(&classes)
	//打印查询出来的数据
	fmt.Println("class>>> ", classes)
	//响应添加学生页面
	ctx.HTML(http.StatusOK, "getStuAdd.html", gin.H{
		"classes": classes,
	})
}

func PostStudentAdd(ctx *gin.Context) {
	//(1). 获取前端页面传入的数据
	//PostForm接收的数据是string类型
	//学号
	strSid := ctx.PostForm("sid")
	//结构体对象中的成员变量是int类型,需要做类型转换
	intSid, _ := strconv.Atoi(strSid)
	//学生姓名
	name := ctx.PostForm("name")
	//学生性别
	strGender := ctx.PostForm("gender")
	//类型转换
	intGender, _ := strconv.Atoi(strGender)
	//学生密码
	pwd := ctx.PostForm("pwd")
	//学生手机号
	tel := ctx.PostForm("tel")
	//所在班级ID
	strClId := ctx.PostForm("clId")
	//类型转换
	intClId, _ := strconv.Atoi(strClId)
	//备注
	remark := ctx.PostForm("remark")
	//(2). 数据验证

	//(3). 添加到数据库
	//实例化结构体对象  性别由int转为byte,使用byte()
	stu := Student{BaseModel: BaseModel{Name: name}, Sno: intSid, Gender: byte(intGender), Pwd: pwd, Tel: tel, ClassID: intClId, Remark: remark}
	//添加到数据库
	DB.Create(&stu)

	//添加完学生后,相应的班级人数也应该加1
	//使用gorm表达式实现
	//先找表,Class{} 这个表,条件Class{}表里的ID等于intClId,然后再写 Class{} 这个表的Num字段再写更新表达式加1
	DB.Model(&Class{}).Where("id = ?", intClId).Update("num", gorm.Expr("num+1"))
	//响应
	//添加成功后重定向到学生管理页面
	//http.StatusMovedPermanently 301 永久重定向 ,302临时重定向,这里使用永久重定向
	ctx.Redirect(http.StatusMovedPermanently, "/student")

}

//删除学生

func GetDeleteStudent(ctx *gin.Context) {
	//获取delID
	delID := ctx.Param("delID")
	//删除符合条件的学生
	DB.Where("sno = ?", delID).Delete(&Student{})

	//响应,重定向到学生管理页面
	ctx.Redirect(http.StatusMovedPermanently, "/student")

}

//编辑学生

func GetEditStudent(ctx *gin.Context) {
	//获取前端传过来的主键ID
	editID := ctx.Param("editID")
	//从数据库查询这个学生的信息
	var editStu Student
	DB.Where("sno = ?", editID).Find(&editStu)

	//从数据库查询班级信息
	var classes []Class
	DB.Find(&classes)
	//响应
	ctx.HTML(http.StatusOK, "getEditStu.html", gin.H{
		"classes": classes,
		"editStu": editStu,
	})
}
