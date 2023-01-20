package model

import "time"

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
	//一对一用户表
	UserInfoID int `gorm:"unique"`
	UserInfo   UserInfo
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
	Teacher Teacher
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

	//一对一 建立学生关联UserInfo表
	UserInfoID int `gorm:"unique"`
	UserInfo   UserInfo
}

//超级管理员模型

type Admin struct {
	BaseModel
	//一对一用户表
	UserInfoID int `gorm:"unique"`
	UserInfo   UserInfo
}
