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
	////查询所有的老师
	////定义[]Teacher
	//var teachers []Teacher
	////查询所有Teacher表中数据写入teachers
	//db.Find(&teachers)
	///*db.Find 相当于 对应的原生SQL
	//先查: select * from teachers;
	//2,liu,2023-02-14 21:49:32.453,2023-02-14 21:49:32.453,2023-02-14 21:49:32.453,1,1001,123,12345678901,,初级讲师
	//3,han,2023-02-14 22:12:29.537,2023-02-14 22:12:29.537,2023-02-14 22:12:29.537,0,1002,123,10987654321,,""
	//在实例化
	//t1 := Teacher{Name:"liu",Tno:1001....}
	//t2 := Teacher{Name:"han",Tno:1002....}
	//放到切片中
	//teachers := []Teacher{t1,t2,t3}
	//*/
	//fmt.Println("teachers: ", teachers)
	//
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"teachers": teachers,
	//})

	//查询单条记录
	////1. 查询单条就不用定义切片了
	//var stu Student
	////先把所有查询出来,然后 按照默认排序取第一条,如果表里的数据量大,就会产生性能问题
	///*[1.061ms] [rows:1] SELECT * FROM `students`
	//stu:  {{1 顾水云 2023-02-15 19:37:26.752 +0800 CST 2023-02-15 19:37:26.752 +0800 CST 2023-02-15 19:37:26.752 +0800 CST} 2001 123 18762531646 0 <nil> 2001级新生 3 {{0  <nil> <nil> <nil>} 0 0 {{0  <nil> <nil> <nil>} 0   0 <nil> }} []}
	//[GIN] 2023/02/17 - 01:24:42 | 200 |    1.823833ms |       127.0.0.1 | GET      "/queryDB"
	//*/
	//db.Find(&stu)
	//
	//fmt.Println("stu: ", stu)
	//
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"stu": stu,
	//})

	////2.查询第一条记录
	//var stu Student
	////数据库查询第一条
	////[0.635ms] [rows:1] SELECT * FROM `students` ORDER BY `students`.`id` LIMIT 1
	//db.First(&stu)
	//
	////打印取出的数据
	//fmt.Println(stu)
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"stu": stu,
	//})

	////3.查询最后一条记录
	//var stu Student
	////数据库查询最后一条
	////[1.019ms] [rows:1] SELECT * FROM `students` ORDER BY `students`.`id` DESC LIMIT 1
	//db.Last(&stu)
	//
	////打印取出的数据
	//fmt.Println(stu)
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"stu": stu,
	//})

	////4.Take一条记录
	//var stu Student
	//
	////[0.770ms] [rows:1] SELECT * FROM `students` LIMIT 1
	//db.Take(&stu)
	//
	////打印取出的数据
	//fmt.Println(stu)
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"stu": stu,
	//})

	////5.条件查询
	//var stu Student
	////where
	////1.查询id为3的学生
	////db.Where("id = ?", 3).Find(&stu)
	////查询学号为200101的学生
	//db.Where("sno = ?", 200103).Find(&stu)
	//
	////打印取出的数据
	//fmt.Println(stu)
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"stu": stu,
	//})

	////6.条件查询
	////单个使用Teacher  多个使用[]Teacher
	//var students []Student
	////where string
	////查询性别为女的学生
	////db.Where("gender = ?", 0).Find(&students)
	////查询性别为男的学生
	//db.Where("gender = ?", 1).Find(&students)
	//
	////打印取出的数据
	//fmt.Println(students)
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"stu": students,
	//})

	//7.一对多的关联关系
	//var student []Student
	//db.Find(&student)
	//fmt.Println("students>>> ", student)

	//定义学生变量
	var student []Student
	//Preload预加载Class这张表 ,预加载的表是跟后边查询的表必须是相关的
	db.Preload("Class").Find(&student)
	fmt.Println("students>>> ", student)

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

//更新

func Update(ctx *gin.Context) {
	//1.不常用 更新所有字段
	////定义学生变量
	//var student Student
	////更新id为3的名字改为乔治
	////先查询
	//db.Where("id = ?", 3).Find(&student)
	//fmt.Println("id为3的学生姓名>>> ", student.Name)
	////重新赋值
	//student.Name = "乔治"
	////保存到数据库中
	//db.Save(&student) //所有字段全部更新
	//fmt.Println("id为3的学生姓名>>> ", student.Name)
	//
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"msg": "更新成功",
	//})

	////1.不常用 更新所有字段
	////更新第一个学生的名字
	////定义学生变量
	//var student Student
	////先查询
	//db.Find(&student)
	////打印学生姓名
	//fmt.Println("姓名: ", student.Name)
	////再更新
	//student.Name = "顾水云"
	////保存到数据库
	//db.Save(&student)
	////打印学生姓名
	//fmt.Println("姓名: ", student.Name)
	//
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"student": student,
	//})

	//2.针对字段的更新 update更新
	//定义学生变量
	//var stu Student //声明未赋值 条件是必须的,这样是不行的 WHERE conditions required
	//var stu = Student{BaseModel: BaseModel{Name: "乔治"}} //还是不行,会提示 WHERE conditions required
	//db.Model 必须是主键信息  乔治的主键ID是3
	var stu = Student{BaseModel: BaseModel{ID: 3}}

	//数据库更新
	//修改乔治的名字变成佩奇
	//UPDATE `students` SET `name`='佩奇' WHERE `id` = 3
	//更新单个字段
	//db.Model(&stu).Update("name", "佩奇")
	//更新多个字段 使用db.Model(&stu).Updates()
	db.Model(&stu).Updates(Student{BaseModel: BaseModel{Name: "查理"}, Remark: "小猪查理烤肉"})

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"stu": stu,
	})

}

//删除

func Delete(ctx *gin.Context) {
	//1.删除单条记录
	////也是先查询,在删除  安装主键删除
	////删除ID为3的查理
	////DELETE FROM `students` WHERE `students`.`id` = 3
	//var stu = Student{BaseModel: BaseModel{ID: 3}}
	////数据库删除
	//db.Delete(&stu)

	//2.删除多条记录
	//按条件删除 删除student表里性别为男的学生
	//DELETE FROM `students` WHERE gender = 1
	db.Where("gender = ?", 1).Delete(&Student{})

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "删除成功!",
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

	//更新
	r.GET("/update", Update)

	//删除
	r.GET("/delete", Delete)

	//启动
	r.Run()
}
