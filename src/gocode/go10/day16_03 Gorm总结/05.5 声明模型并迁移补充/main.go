package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

//基本模型类

type BaseModel struct {
	ID         int        `gorm:"primaryKey"`
	Name       string     `gorm:"type:varchar(32);unique;not null"`
	CreateTime *time.Time `gorm:"autoCreateTime"`
	UpdateTime *time.Time `gorm:"autoCreateTime"`
	DeleteTime *time.Time `gorm:"autoCreateTime"`
}

//教师类

type Teacher struct {
	//继承BaseModel模型类
	BaseModel
	//性别
	Gender byte `gorm:"default:0"`
	//教师编号
	Tno int
	//密码
	Pwd string `gorm:"type:varchar(100);not null"`
	//手机号
	Tel string `gorm:"type:char(11);"`
	//生日
	Birth *time.Time
	//备注 初级讲师 中级讲师 高级讲师
	Remark string `gorm:"type:varchar(255);"`
}

//班级类

type Class struct {
	//继承BaseModel模型类
	BaseModel
	//班级人数
	Num int

	//一对多
	//一个班级对应一个导员教师
	//导员ID
	TutorID int
	//关联教师表
	Tutor Teacher
}

//课程类

type Course struct {
	//继承BaseModel模型类
	BaseModel
	//课程学分
	Credit int
	//课程周期
	Period int

	//多对一
	//多个课程,可以有一个老师教学
	TeacherID int
	//关联教师表 外键关联
	//Teacher Teacher `gorm:"foreignKey:TeacherID"`
	//如果不写`gorm:"foreignKey:TeacherID"` 要写TeacherID, 默认是Teacher Teacher的第一个拼上ID
	Teacher Teacher
}

//学生类

type Student struct {
	//继承BaseModel模型类
	BaseModel
	//学号
	Sno int
	//性别
	Gender byte `gorm:"default:0"`
	//密码
	Pwd string `gorm:"type:varchar(100);not null"`
	//手机号
	Tel int `gorm:"type:char(11);"`
	//生日
	Birth *time.Time
	//备注
	Remark string `gorm:"type:varchar(255);"`

	//多对一
	//一个班级可以有多个学生
	ClassID int
	//关联班级表
	Class Class

	//多对多
	//一个学生可以选择多个课程,一个课程业可以被多个学生选择
	//创建第三张关联关系表,叫student2course ,级联删除
	Courses []Course `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`
}

//以上的模型类会创建五张表

func main() {

	//定义数据库连接信息 dsn
	dsn := "root:12345678@tcp(127.0.0.1:3306)/css?charset=utf8mb4&parseTime=True&loc=Local"

	//创建日志对象
	nerLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io write
		logger.Config{
			LogLevel: logger.Info, //日志级别 Log lever
		},
	)

	//GORM连接配置
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: nerLogger, //日志配置
	})
	//处理错误
	if err != nil {
		panic("Filed to connect database...")
	}
	//自动迁移
	db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&Class{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Student{})

}
