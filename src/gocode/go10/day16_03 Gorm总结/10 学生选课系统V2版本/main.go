package main

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//基本模型类

type BaseModel struct {
	//ID 主键
	ID int `gorm:"primaryKey"`
	//姓名 string类型
	Name string `gorm:"type:varchar(32);unique;not null"`
	//创建时间
	CreateTime *time.Time `gorm:"autoCreateTime"`
	//修改时间
	UpdateTime *time.Time `gorm:"autoCreateTime"`
	//删除时间
	DeleteTime *time.Time `gorm:"autoCreateTime"`
}

//Teacher  教师表

type Teacher struct {
	//继承基本模型表
	BaseModel
	//教师编号
	Tno int
	//密码
	Pwd string `gorm:"type:varchar(100);not null"`
	//手机号
	Tel string `gorm:"type:char(11);"`
	//性别
	Gender byte `gorm:"default:0"`
	//生日
	Birth *time.Time
	//备注 (初级讲师  中级讲师  高级讲师)
	Remark string `gorm:"type:varchar(255);"`
}

//Class  班级表

type Class struct {
	//继承基本模型表
	BaseModel
	//人数
	Num int
	//TutorID
	TutorID int
	//关联教师表
	Tutor Teacher
}

//Course 课程表

type Course struct {
	//继承基本模型表
	BaseModel
	//学分
	Credit int
	//周期
	Period int

	//多对一
	//对应 教师ID
	TeacherID int
	//关联教师表
	Teacher Teacher
}

//Student 学生表

type Student struct {
	//继承基本模型表
	BaseModel
	//学号
	Sno int
	//密码
	Pwd string `gorm:"type:varchar(100);not null"`
	//手机号
	Tel string `gorm:"type:char(11);"`
	//性别
	Gender byte `gorm:"default:0"`
	//生日
	Birth *time.Time
	//备注
	Remark string `gorm:"type:varchar(255)"`

	//多对一
	//课程ID
	ClassID int
	//外检约束 Class关联班级表,并外键约束ClassID
	Class Class `gorm:"foreignKey:ClassID"`
	//多对多
	//会自动创建第三张关系表student2course
	Courses []Course `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`
}

//数据库连接函数

func DBInit() *gorm.DB {
	//定义dsn 数据库连接信息
	dsn := "root:12345678@tcp(127.0.0.1:3306)/css?charset=utf8mb4&parseTime=True&loc=Local"

	//创建日志对象
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io  write
		logger.Config{
			LogLevel: logger.Info, // Log level(日志级别)
		},
	)

	//获取db对象 ,db对象类似文件句柄
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//日志配置
		Logger: newLogger,
	})

	//判断错误
	if err != nil {
		panic("Failed to connect to database ...")
	}

	fmt.Println("数据库连接成功!!!", db)

	//自动迁移
	db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&Class{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Student{})

	return db
}

// 数据库初始化
var db = DBInit()

// 首页路由函数
func getIndex(ctx *gin.Context) {
	//响应首页页面
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// 学生路由函数
func getStudent(ctx *gin.Context) {
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
		db.Preload("Class.Tutor").Find(&student)
		//响应学生页面
		ctx.HTML(http.StatusOK, "student.html", gin.H{
			"student": student,
		})
	} else {
		//数据库查询
		//模糊查询 "name like ?", "%"+name+"%"  包含name参数的
		db.Where("name like ?", "%"+name+"%").Preload("Class.Tutor").Find(&student)
		//响应学生页面
		ctx.HTML(http.StatusOK, "student.html", gin.H{
			"student": student,
		})
	}

}

//查看学生信息详情

func getOneStuHTml(ctx *gin.Context) {
	//获取动态参数
	strSno := ctx.Param("sno")
	//强转
	intSno, _ := strconv.Atoi(strSno)

	//定义学生信息
	var stu Student
	//数据库查询 预加载班级表
	db.Where("sno = ? ", intSno).Preload("Class").Take(&stu)

	////响应
	//ctx.JSON(200, gin.H{
	//	"intSno": stu,
	//})

	//响应学生个人详情页面
	ctx.HTML(http.StatusOK, "detailStudent.html", gin.H{
		"stu": stu,
	})

	//fmt.Println("stu >>> ", stu)

}

// 学生选课
//获取选课页面

func getSelectCourseHtml(ctx *gin.Context) {
	//获取学生学号
	strSno := ctx.Param("sno")
	intSno, _ := strconv.Atoi(strSno)

	//数据库查询
	var stu Student
	db.Where("sno = ?", intSno).Preload("Class").Take(&stu)

	//获取所有课程
	var courses []Course
	db.Preload("Teacher").Find(&courses)

	//响应选课页面
	ctx.HTML(http.StatusOK, "selectCourse.html", gin.H{
		"stu":     stu,
		"courses": courses,
	})
}

//提交选课

func postSelectCourse(ctx *gin.Context) {
	//获取前端传来的数据 获取多个值,键值对使用PostFormArray
	courseID := ctx.PostFormArray("course_id")
	fmt.Println(courseID)

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"courseID": courseID,
	})
}

//添加学生路由函数

func GetStudentAddHtml(ctx *gin.Context) {
	//查询班级信息
	var classes []Class
	//数据库查询
	db.Find(&classes)
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
	db.Create(&stu)

	//添加完学生后,相应的班级人数也应该加1
	//使用gorm表达式实现
	//先找表,Class{} 这个表,条件Class{}表里的ID等于intClId,然后再写 Class{} 这个表的Num字段再写更新表达式加1
	db.Model(&Class{}).Where("id = ?", intClId).Update("num", gorm.Expr("num+1"))
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
	db.Where("sno = ?", delID).Delete(&Student{})

	//响应,重定向到学生管理页面
	ctx.Redirect(http.StatusMovedPermanently, "/student")

}

//编辑学生

func GetEditStudent(ctx *gin.Context) {
	//获取前端传过来的主键ID
	editID := ctx.Param("editID")
	//从数据库查询这个学生的信息
	var editStu Student
	db.Where("sno = ?", editID).Find(&editStu)

	//从数据库查询班级信息
	var classes []Class
	db.Find(&classes)
	//响应
	ctx.HTML(http.StatusOK, "getEditStu.html", gin.H{
		"classes": classes,
		"editStu": editStu,
	})
}

// 课程路由函数
func getCourse(ctx *gin.Context) {
	//获取课程名
	coName := ctx.Query("coName")
	//从数据库查询所有课程
	var courses []Course

	//基于模糊查询的搜索框
	if coName == "" {
		//查询
		db.Find(&courses)

		//响应 查出的数据传给前端
		ctx.HTML(http.StatusOK, "course.html", gin.H{
			"courses": courses,
		})
	} else {
		//查询
		db.Where("name like ?", "%"+coName+"%").Find(&courses)
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
	db.Find(&teachers)
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
	db.Create(&course)

	//响应,重定向到课程管理页面
	ctx.Redirect(http.StatusMovedPermanently, "/course")

}

//删除课程

func GetDeleteCourse(ctx *gin.Context) {
	//获取delID
	delID := ctx.Param("delID")
	//先查询,在删除
	db.Where("ID = ?", delID).Delete(&Course{})

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
	db.Where("ID = ?", editID).Find(&course)

	//查询teacher表
	var teachers []Teacher
	db.Find(&teachers)

	//响应编辑页面
	ctx.HTML(http.StatusOK, "editCourse.html", gin.H{
		"course":   course,
		"teachers": teachers,
	})

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
	//查看学生的个人信息
	r.GET("/student/:sno", getOneStuHTml)
	//学生选课页面
	r.GET("/student/selectCourse/:sno", getSelectCourseHtml)
	//学生选课提交
	r.POST("/student/selectCourse/:sno", postSelectCourse)
	//获取添加学生页面
	r.GET("/student/add", GetStudentAddHtml)
	//添加学生
	r.POST("/student/add", PostStudentAdd)
	//删除学生
	r.GET("/student/delete/:delID", GetDeleteStudent)
	//编辑学生
	r.GET("/student/edit/:editID", GetEditStudent)

	//课程路由
	r.GET("/course", getCourse)
	//添加课程页面 Get 拿页面  Post提交课程
	r.GET("/course/add", GetCourseAddHtml)
	//添加课程
	r.POST("/course/add", PostCourseAdd)
	//删除课程
	r.GET("/course/delete/:delID", GetDeleteCourse)
	//编辑课程
	r.GET("/course/edit/:editID", GetEditCourseHtml)
	//提交课程

	//班级路由
	r.GET("/class", getClass)

	//启动
	r.Run()

}
