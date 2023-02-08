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
	dsn := "dev:xxxxx@tcp(xxxxxxx)/css?charset=utf8mb4&parseTime=True&loc=Local"

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

//添加

func add(ctx *gin.Context) {
	//多对多的添加

	////实例化结构体对象
	//stu := Student{
	//	BaseModel: BaseModel{Name: "宋晶爱"},
	//	Sno:       202201,
	//	Pwd:       "123",
	//	Tel:       "13641166790",
	//	Gender:    0,
	//	Remark:    "2022级新生",
	//	ClassID:   2,
	//}
	//fmt.Println("before create stu ID", stu.ID)
	////添加到数据库中
	//db.Create(&stu)
	//fmt.Println("after create stu ID", stu.ID)

	//多对多的关系绑定 [学生和课程]
	//为学生绑定课程
	//var courses []Course
	//db.Where("id in ?", []int{1, 2}).Find(&courses)
	//fmt.Println("courses: ", courses)

	////1.创建学生的同时绑定课程
	//stu := Student{
	//	BaseModel: BaseModel{Name: "刘群秀"},
	//	Sno:       202202,
	//	Pwd:       "123",
	//	Tel:       "18480292668",
	//	Gender:    0,
	//	Remark:    "2022级新生",
	//	ClassID:   2,
	//	Courses:   courses,
	//}
	////添加到数据库
	//db.Create(&stu)

	//2.先创建学生,,在绑定课程
	//为李四绑定数据库和数据结构两门课程
	//var courses []Course
	//db.Where("name in ?", []string{"数据库设计", "数据结构"}).Find(&courses)
	//fmt.Println("courses: ", courses)
	////查找学生  查找李四
	//var stu Student
	//db.Where("name = ?", "李四").Take(&stu)
	//给李四学生stu绑定courses课程
	//append添加
	//db.Model(&stu).Association("Courses").Append(courses)
	//clear()里边不加参数 取消关联
	//db.Model(&stu).Association("Courses").Clear()

	//find()查询
	//查找学生  查找李四
	var stu Student
	db.Where("name = ?", "李四").Take(&stu)
	//创建课程对象
	var courses []Course
	//关联课程表
	db.Model(&stu).Association("Courses").Find(&courses)

	//响应
	ctx.JSON(200, gin.H{
		"msg":     "success",
		"courses": courses,
	})
}

//查询

func selects(ctx *gin.Context) {

	//一对多的关联查询(学生和班级 多对一的关系)
	var student Student

	//查询所有(查表)
	//Preload预加载
	db.Preload("Class.Tutor").Find(&student)
	fmt.Println("Student::: ", student)

	ctx.JSON(200, gin.H{
		"msg":     "success",
		"Student": student,
	})
}

//更新

func update(ctx *gin.Context) {
	////一: 不常用:更新所有字段
	//var student Student
	//db.First(&student)
	//fmt.Println(student.ID)
	//fmt.Println(student.Name)
	////student.Name = "张三三"  并不会修改数据库中的值,只是修改的结构体对象中的值
	////同步数据库
	//student.Name = "张三三"
	//db.Save(&student) //所有字段全部更新
	////UPDATE `students` SET `name`='张三三',`create_time`='2022-12-09 23:20:40.366',`update_time`='2022-12-09 23:20:40.366',`delete_time`='2022-12-09 23:20:40.366',`sno`=200201,`pwd`='123456',`tel`='12345678901',`gender`=1,`birth`=NULL,`remark`='新生',`class_id`=2 WHERE `id` = 2

	//二 针对字段更新 update方法
	//var s Student //声明未赋值  WHERE conditions required
	//var s = Student{BaseModel: BaseModel{Name: "李四"}}  //依旧是WHERE conditions required
	var s = Student{BaseModel: BaseModel{ID: 5}}
	//model查询   update修改
	//db.Model(&s) 里的必须是主键信息
	//db.Model(&s).Update("name", "李思思") //修改一个字段
	db.Model(&s).Updates(Student{BaseModel: BaseModel{Name: "刘琪"}, Gender: '0'}) //修改多个字段

	ctx.JSON(200, gin.H{
		"msg": s,
	})
}

//删除

func delete(ctx *gin.Context) {
	////删除单条记录
	////实例化删除对象
	//var s = Student{BaseModel: BaseModel{ID: 2}}
	////删除表记录
	//db.Delete(&s)

	////按照ID删除
	//var s = Student{BaseModel: BaseModel{ID: 3}}
	////删除表记录
	//db.Delete(&s)

	//删除多条记录
	//按条件删除,删除gender=1的记录
	//db.Where("gender = ?", 1).Delete(Student{})

	db.Where("gender = ?", 2).Delete(Student{})

	ctx.JSON(200, gin.H{
		"msg": "删除成功!",
	})
}

func main() {

	//创建路由对象
	r := gin.Default()

	////数据库初始化
	//DBInit()

	fmt.Println(r)

	//添加
	r.POST("/add", add)

	////查询
	//r.GET("/select", listStu)
	//查询
	r.GET("/selects", selects)
	//更新
	r.GET("update", update)
	//删除
	r.GET("/delete", delete)

	//启动
	r.Run()
}
