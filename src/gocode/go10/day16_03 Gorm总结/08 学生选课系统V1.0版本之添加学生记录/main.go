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

	//响应学生页面
	ctx.HTML(http.StatusOK, "student.html", gin.H{})
}

//添加学生路由函数

func GetStudentAddHtml(ctx *gin.Context) {
	//响应添加学生页面
	ctx.HTML(http.StatusOK, "getStuAdd.html", nil)
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

	//响应
	//添加成功后重定向到学生管理页面
	//http.StatusMovedPermanently 301 永久重定向 ,302临时重定向,这里使用永久重定向
	ctx.Redirect(http.StatusMovedPermanently, "/student")

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
	//添加学生
	render.AddFromFiles("getStuAdd.html", "templates/base.html", "templates/getStuAdd.html")
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
	//获取添加学生页面
	r.GET("/student/add", GetStudentAddHtml)
	//添加学生
	r.POST("/student/add", PostStudentAdd)
	//课程路由
	r.GET("/course", getCourse)
	//班级路由
	r.GET("/class", getClass)

	//启动
	r.Run()

}
