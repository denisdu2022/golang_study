package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
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

//添加

func Add(ctx *gin.Context) {
	//添加记录
	//一个结构体对象映射一条表记录
	//添加单条记录
	//Gender 0为女 1为男
	//t1 := Teacher{BaseModel: BaseModel{Name: "liu"}, Tno: 1001, Pwd: "123", Tel: "12345678901", Gender: 1, Remark: "初级讲师"}
	////打印实例化对象数据
	//fmt.Println(t1)
	//fmt.Println(t1.Name)
	////添加单条记录到数据库中
	//db.Create(&t1) //翻译成insert语句
	////插入到数据库完成后 GORM 将生成一条SQL语句来插入数据并回填主键值
	//fmt.Println(t1.Name)
	//fmt.Println(t1.ID)

	////补充
	////通过前端的内容设置数据库
	////获取前端form表单数据
	////获取name
	//name := ctx.PostForm("name")
	////获取tno
	//tno := ctx.PostForm("tno")
	////结构体的成员变量类型是int   强转为int
	//iTno, _ := strconv.Atoi(tno)
	////获取密码
	//pwd := ctx.PostForm("pwd")
	////获取tel
	//tel := ctx.PostForm("tel")
	////获取性别
	//gender := ctx.PostForm("gender")
	////强转为int
	//iGender, _ := strconv.Atoi(gender)
	////Gender 结构体的成员变量类型是byte()    使用byte(iGender)
	//t1 := Teacher{BaseModel: BaseModel{Name: name}, Tno: iTno, Pwd: pwd, Tel: tel, Gender: byte(iGender)}
	//
	////添加到数据库记录
	//db.Create(&t1)

	////2.创建多条记录
	//t1 := Teacher{BaseModel: BaseModel{Name: "sun"}, Tno: 1003, Pwd: "123", Tel: "111", Gender: 0, Remark: "初级讲师"}
	//t2 := Teacher{BaseModel: BaseModel{Name: "zhao"}, Tno: 1004, Pwd: "123", Tel: "111", Gender: 0, Remark: "初级讲师"}
	//t3 := Teacher{BaseModel: BaseModel{Name: "qian"}, Tno: 1005, Pwd: "123", Tel: "111", Gender: 0, Remark: "初级讲师"}
	//t4 := Teacher{BaseModel: BaseModel{Name: "li"}, Tno: 1006, Pwd: "123", Tel: "111", Gender: 0, Remark: "初级讲师"}
	//
	////批量添加记录时创建一个切片
	//teachers := []Teacher{t1, t2, t3, t4}
	////添加到数据库记录
	//db.Create(&teachers)
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"teachers": teachers,
	//	"msg":      "添加成功",
	//})

	////3.创建一对多的记录
	//cl1 := Class{BaseModel: BaseModel{Name: "软件一班"}, Num: 36, TutorID: 2}
	//
	////添加到数据库记录
	//db.Create(&cl1)

	//批量创建一对多的记录
	cl1 := Class{BaseModel: BaseModel{Name: "软件一班"}, Num: 36, TutorID: 2}
	cl2 := Class{BaseModel: BaseModel{Name: "软件二班"}, Num: 38, TutorID: 8}
	cl3 := Class{BaseModel: BaseModel{Name: "计算机科学与技术一班"}, Num: 38, TutorID: 3}
	cl4 := Class{BaseModel: BaseModel{Name: "计算机科学与技术二班"}, Num: 38, TutorID: 9}
	//批量添加创建切片
	classes := []Class{cl1, cl2, cl3, cl4}
	//添加到数据库记录
	db.Create(&classes)

	//打印班级ID
	for _, class := range classes {
		fmt.Println(class.ID) //班级ID 3,4,5,6
	}

	////指定批量添加的大小
	//db.CreateInBatches(classes,100)
	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "添加成功!",
		"cl":  classes,
	})

}

//查询

func QueryDB(ctx *gin.Context) {
	//查询所有的老师
	//定义[]Teacher
	var teachers []Teacher
	//查询所有Teacher表中数据写入teachers
	db.Find(&teachers)
	/*db.Find 相当于 对应的原生SQL
	先查: select * from teachers;
	2,liu,2023-02-14 21:49:32.453,2023-02-14 21:49:32.453,2023-02-14 21:49:32.453,1,1001,123,12345678901,,初级讲师
	3,han,2023-02-14 22:12:29.537,2023-02-14 22:12:29.537,2023-02-14 22:12:29.537,0,1002,123,10987654321,,""
	在实例化
	t1 := Teacher{Name:"liu",Tno:1001....}
	t2 := Teacher{Name:"han",Tno:1002....}
	放到切片中
	teachers := []Teacher{t1,t2,t3}
	*/
	fmt.Println("teachers: ", teachers)

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"teachers": teachers,
	})
}

func main() {

	//初始化路由对象
	r := gin.Default()

	//路由
	//添加
	r.POST("/add", Add)

	//查询
	r.GET("/queryDB", QueryDB)

	//启动
	r.Run()
}
