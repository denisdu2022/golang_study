package main

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"time"
)

//公共的模型表

type BaseModel struct {
	ID         int        `gorm:"primaryKey"`
	Name       string     `gorm:"type:varchar(32);unique;not null"`
	CreateTime *time.Time `gorm:"autoCreateTime"` //创建时间
	UpdateTime *time.Time `gorm:"autoCreateTime"` //更新时间
	DeleteTime *time.Time `gorm:"autoCreateTime"` //删除时间
}

// 教师表

type Teacher struct {
	BaseModel
	Ton    int
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:char(11);"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255);"`
}

//课程表
//Teacher 就是TeacherID

type Course struct {
	BaseModel
	Credit int //学分
	Period int //周期
	//一对多
	TeacherID int
	//写全
	//Teacher Teacher
	//缩写
	Teacher
}

//班级表

type Class struct {
	BaseModel
	Num     int //班级人数
	TutorID int
	Tutor   Teacher
}

//学生表

type Student struct {
	BaseModel
	Sno    int
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:char(11);"`
	Gender byte   `gorm:"default:1"`
	Birth  *time.Time
	Remark string `gorm:"type:varchar(255);"`
	//建立一对多的关联关系
	ClassID int
	//用那张表的主键去关联那个字段  外键约束
	Class Class `gorm:"foreignKey:ClassID"`

	//建立多对多的关联关系:通过创建第三张表 courses 只是orm层 结构体的成员变量,在sql层没有这个字段,会有一张多对多的关系表
	Courses []Course `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`
}

//数据库初始化函数

// 数据库初始化
var db *gorm.DB

func DBInit() *gorm.DB {
	//数据库初始化
	//数据库连接信息
	//dsn := "root:go20222023@tcp(127.0.0.1:3306)/css?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "dev:CIAy2k0308@tcp(mysql-internet-cn-north-1-3156338edfb24ab1.rds.jdcloud.com:3306)/css?charset=utf8mb4&parseTime=True&loc=Local"

	//创建日志对象
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //io writer
		logger.Config{
			//SlowThreshold: time.Second, //慢SQL阈值
			LogLevel: logger.Info, //Log lever(日志级别)
		},
	)

	//连接打开数据库 gorm.Open连接数据库  &gorm.Config{}是配置使用
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//日志配置
		Logger: newLogger,
	})
	//db对象类似文件句柄,后续对数据库操作都使用db对象
	fmt.Println(db)

	//迁移模型类: 将模型类转换为SQL表
	//db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Class{})
	db.AutoMigrate(&Student{})
	return db
}

//// 结构体
//
//type Student struct {
//	Sid    int
//	Name   string
//	Age    int
//	Gender string
//	Email  string
//	Addr   string
//}

func index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func student(ctx *gin.Context) {
	//获取过滤参数
	user := ctx.Query("user")
	fmt.Println("user:::", user)

	//查询student表中的记录
	var students []Student
	//判断是否携带参数
	if user == "" {
		db.Find(&students)
	} else {
		db.Where("name like ?", "%"+user+"%").Find(&students)
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
	render.AddFromFiles("getStuAdd", "templates/base.html", "templates/getStuAdd.html")
	render.AddFromFiles("class", "templates/base.html", "templates/class.html")
	render.AddFromFiles("course", "templates/base.html", "templates/course.html")

	return render
}

//添加学生

//

//Get请求

func GetStuAdd(ctx *gin.Context) {
	ctx.HTML(200, "getStuAdd", nil)
}

//Post请求

func PostStuAdd(ctx *gin.Context) {
	//1.获取前端请求数据
	var sid, genderInt, classID int
	//转换为int
	sid, _ = strconv.Atoi(ctx.PostForm("sid"))
	name := ctx.PostForm("user")
	pwd := ctx.PostForm("pwd")
	tel := ctx.PostForm("tel")
	genderInt, _ = strconv.Atoi(ctx.PostForm("gender"))
	//强转为uint8
	gender := uint8(genderInt)
	//班级
	classID, _ = strconv.Atoi(ctx.PostForm("class"))
	remark := ctx.PostForm("remark")
	fmt.Println(sid, name, pwd, tel, genderInt, gender, classID, remark)
	//2.数据验证

	//3.添加数据到数据库存储
	//实例化结构体对象
	addStudent := Student{BaseModel: BaseModel{Name: name}, Sno: sid, Pwd: pwd, Tel: tel, Gender: gender, ClassID: classID, Remark: remark}
	//插入到数据库
	db.Create(&addStudent)

	//4.响应
	ctx.JSON(200, gin.H{
		"stu": addStudent,
		"msg": "学生添加成功!!!",
	})

}

func main() {
	//获取路由对象
	r := gin.Default()
	////加载模版文件
	//r.LoadHTMLGlob("templates/*")

	//初始化数据库
	DBInit()

	//加载模版文件
	r.HTMLRender = createMyRender()

	//设置静态文件资源窗口
	r.Static("/static", "./static")

	//首页
	r.GET("/", index)
	//学生首页页面
	r.GET("/student", student)
	//添加学生
	r.GET("/student/add", GetStuAdd)
	r.POST("/student/add", PostStuAdd)

	//班级页面
	r.GET("/class", class)
	//课程管理
	r.GET("/course", course)

	//启动
	r.Run(":8090")

}
