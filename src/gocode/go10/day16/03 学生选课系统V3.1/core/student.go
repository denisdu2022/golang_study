package core

import (
	. "css/databse"
	. "css/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

//获取所有学生信息

func GetAllStudent(ctx *gin.Context) {
	//获取过滤参数
	user := ctx.Query("user")
	fmt.Println("user:::", user)

	//查询student表中的记录
	var students []Student
	//判断是否携带参数
	if user == "" {
		//Preload预加载
		DB.Preload("Class.Tutor").Find(&students)
	} else {
		DB.Where("name like ?", "%"+user+"%").Find(&students)
	}

	ctx.HTML(200, "student.html", gin.H{
		"students": students,
	})
}

//获取某一个学生个人信息详情

func GetOneStudent(ctx *gin.Context) {
	//获取学生学号
	sno := ctx.Param("sno")
	//获取学生对象
	var student Student
	//单条查询,并预加载班级表,以获得班级名称
	DB.Where("sno = ? ", sno).Preload("Class").Take(&student)
	//获取当前学生的所选课程
	var selectCourse []Course
	//预加载teacher表
	DB.Preload("Teacher").Model(&student).Association("Courses").Find(&selectCourse)

	ctx.HTML(200, "detailOneStudent.html", gin.H{
		"student":      student,
		"selectCourse": selectCourse,
	})
}

// 学生个人选课页面
func GetSelectCourse(ctx *gin.Context) {
	//获取学生学号
	sno := ctx.Param("sno")
	//获取学生对象
	var student Student
	//单条查询,并预加载班级表,以获得班级名称
	DB.Where("sno = ? ", sno).Preload("Class").Take(&student)
	//获取所有课程
	var courses []Course
	//查询课程
	DB.Preload("Teacher").Find(&courses)

	ctx.HTML(200, "selectCourse.html", gin.H{
		"student": student,
		"courses": courses,
	})
}

// 学生个人选课

func SelectCourse(ctx *gin.Context) {
	//获取学生学号
	sno := ctx.Param("sno")
	//获取学生对象
	var student Student
	//单条查询,并预加载班级表,以获得班级名称
	DB.Where("sno = ? ", sno).Preload("Class").Take(&student)
	//获取前端的course_id
	courseIDSlice := ctx.PostFormArray("course_id")
	//获取所有课程
	var courses []Course
	//查询课程
	DB.Where("id in ?", courseIDSlice).Find(&courses)
	//为学号为sno的学生对象,绑定课程courses
	DB.Model(&student).Association("Courses").Append(&courses)

	//ctx.HTML(200, "selectCourse.html", gin.H{
	//	"student": student,
	//	"courses": courses,
	//})
	ctx.Redirect(301, "/student/"+sno)
}

//添加学生

//Get请求

func GetStuAdd(ctx *gin.Context) {
	//查询所有的班级记录
	var classes []Class
	DB.Find(&classes)
	fmt.Println(classes)
	ctx.HTML(200, "getStuAdd", gin.H{
		"classes": classes,
	})
}

//Post请求

func PostStuAdd(ctx *gin.Context) {
	//1.获取前端请求数据
	//转换为int
	sid, _ := strconv.Atoi(ctx.PostForm("sid"))
	name := ctx.PostForm("user")
	pwd := ctx.PostForm("pwd")
	tel := ctx.PostForm("tel")
	genderInt, _ := strconv.Atoi(ctx.PostForm("gender"))
	//fmt.Println(genderInt)
	//强转为uint8
	gender := uint8(genderInt)
	//fmt.Println(gender)
	//班级
	classID, _ := strconv.Atoi(ctx.PostForm("class"))
	remark := ctx.PostForm("remark")
	fmt.Println("--------------------", sid, name, pwd, tel, genderInt, gender, classID, remark)
	//2.数据验证

	//3.添加数据到数据库存储
	//实例化结构体对象
	addStudent := Student{BaseModel: BaseModel{Name: name}, Sno: sid, Pwd: pwd, Tel: tel, Gender: gender, ClassID: classID, Remark: remark}
	fmt.Println(addStudent)
	//插入到数据库
	DB.Create(&addStudent)

	//6.添加完学生后,班级人数自增1 使用gorm表达式
	//如果是翻译成SQL语句则是 update Class set num = num + 1 where......
	DB.Model(&Class{}).Where("id = ?", classID).Update("num", gorm.Expr("num + 1"))

	//4.响应
	//ctx.JSON(200, gin.H{
	//	"stu": addStudent,
	//	"msg": "学生添加成功!!!",
	//})

	//5.响应 重定向
	//站外重定向
	//ctx.Redirect(http.StatusMovedPermanently, "https://www.cdn-vod.com")
	//站内重定向
	ctx.Redirect(http.StatusMovedPermanently, "/student")

	////如果不重定向,就想看到新的学生页面
	////查询student表中的记录
	//var students []Student
	//ctx.HTML(200, "student.html", gin.H{
	//	"students": students,
	//})

}

//删除学生

func DeleteStu(ctx *gin.Context) {

	//获取delID
	delID := ctx.Param("delID")
	//删除符合条件的学生
	DB.Where("sno = ?", delID).Delete(Student{})

	//响应
	ctx.JSON(200, gin.H{
		"msg": "删除成功",
	})
}

//编辑学生

func GetStuEditHtml(ctx *gin.Context) {
	//获取editID 获取路径参数,即学号
	editID := ctx.Param("editID")
	//查询符合条件的学生对象
	var editStu Student
	DB.Where("sno = ?", editID).Find(&editStu)
	fmt.Println("editStu: ", editStu)

	//查询所有的班级记录
	var classes []Class
	DB.Find(&classes)

	//响应
	ctx.HTML(200, "editStudent.html", gin.H{
		"editStu": editStu,
		"classes": classes,
	})
}
