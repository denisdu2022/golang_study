package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//日志配置
		Logger: newLogger,
	})
	//db对象类似文件句柄,后续对数据库操作都使用db对象
	fmt.Println(db, err)

	//迁移模型类: 将模型类转换为SQL表
	//db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Class{})
	db.AutoMigrate(&Student{})
	return db
}

// 数据库初始化
var db = DBInit()

func add(ctx *gin.Context) {

	//单表
	//ORM映射 结构体对象对应表记录
	//t1 := Teacher{BaseModel: BaseModel{Name: "luobo"}, Ton: 1001, Pwd: "luobo", Tel: "13456767890", Remark: "高级讲师"}
	//fmt.Println(t1)
	//fmt.Println(t1.Name)

	//通过前端的内容设置数据库记录
	//name := ctx.PostForm("name")
	//ton, _ := strconv.Atoi(ctx.PostForm("Ton"))
	//pwd := ctx.PostForm("pwd")
	//tel := ctx.PostForm("tel")
	//remark := ctx.PostForm("remark")
	//t := Teacher{BaseModel: BaseModel{Name: name}, Ton: ton, Pwd: pwd, Tel: tel, Remark: remark}
	//fmt.Println(t)

	////添加单条记录
	//翻译成insert语句
	//db.Create(&t1) //传地址,&,会把插入的数据返回回来

	//前端传值
	//db.Create(&t)

	//ctx.JSON(200, gin.H{
	//	"msg": "数据库表记录添加成功!!!",
	//	//"t1":  t1,
	//	"t": t,
	//})

	//批量添加记录
	//SQL语句中
	//insert table_name values()
	//t1 := Teacher{BaseModel: BaseModel{Name: "luobo1"}, Ton: 1001, Pwd: "luobo1", Tel: "13456767890", Remark: "高级讲师"}
	//t2 := Teacher{BaseModel: BaseModel{Name: "luobo2"}, Ton: 1001, Pwd: "luobo2", Tel: "13456767890", Remark: "高级讲师"}
	//t3 := Teacher{BaseModel: BaseModel{Name: "luobo3"}, Ton: 1001, Pwd: "luobo3", Tel: "13456767890", Remark: "高级讲师"}

	//t1 := Teacher{BaseModel: BaseModel{Name: "luobo"}, Ton: 1001, Pwd: "luobo", Tel: "13456767890", Remark: "高级讲师"}
	//t2 := Teacher{BaseModel: BaseModel{Name: "jack"}, Ton: 1002, Pwd: "jack", Tel: "13456763450", Remark: "中级讲师"}
	//t3 := Teacher{BaseModel: BaseModel{Name: "hah"}, Ton: 1003, Pwd: "hah", Tel: "13678767890", Remark: "初级讲师"}
	////创建切片并赋值
	//teachers := []Teacher{t1, t2, t3}
	////批量添加
	//db.Create(&teachers)
	//
	//ctx.JSON(200, gin.H{
	//	"msg": "数据库表记录添加成功!!!",
	//	//"t1":  t1,
	//	"t": teachers,
	//})

	////创建一对多记录
	//c1 := Class{BaseModel: BaseModel{Name: "软件一班"}, Num: 36, TutorID: 7}
	//db.Create(&c1)

	////创建一对多记录,批量添加
	c1 := Class{BaseModel: BaseModel{Name: "软件一班"}, Num: 36, TutorID: 7}
	c2 := Class{BaseModel: BaseModel{Name: "软件二班"}, Num: 36, TutorID: 7}
	c3 := Class{BaseModel: BaseModel{Name: "计算机科学与技术一班"}, Num: 38, TutorID: 8}
	c4 := Class{BaseModel: BaseModel{Name: "计算机科学与技术2班"}, Num: 38, TutorID: 9}
	classes := []Class{c1, c2, c3, c4}
	db.Create(&classes)

	ctx.JSON(200, gin.H{
		"msg": "数据库表记录添加成功!!!",
		//"t1":  t1,
		"classes": classes,
	})
}

func main() {

	//创建路由对象
	r := gin.Default()

	////数据库初始化
	//DBInit()

	fmt.Println(r)

	r.POST("/add", add)

	//启动
	r.Run()
}
