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

//Class  班级表模型

//Class表添加成员变量Student

type Class struct {
	//继承基本模型表
	BaseModel
	//人数
	Num int
	//TutorID
	TutorID int
	//关联教师表
	Tutor Teacher

	//一对多关系
	Student []Student //反向查询字段,因为有多个学生,所以是一个切片[]Student
}

//Course 课程表模型

//Course表添加成员变量student

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

	//多对多
	Students []Student `gorm:"many2many:student2course;"`
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

	////批量创建一对多的记录
	//cl1 := Class{BaseModel: BaseModel{Name: "软件一班"}, Num: 36, TutorID: 2}
	//cl2 := Class{BaseModel: BaseModel{Name: "软件二班"}, Num: 38, TutorID: 8}
	//cl3 := Class{BaseModel: BaseModel{Name: "计算机科学与技术一班"}, Num: 38, TutorID: 3}
	//cl4 := Class{BaseModel: BaseModel{Name: "计算机科学与技术二班"}, Num: 38, TutorID: 9}
	////批量添加创建切片
	//classes := []Class{cl1, cl2, cl3, cl4}
	////添加到数据库记录
	//db.Create(&classes)
	//
	////打印班级ID
	//for _, class := range classes {
	//	fmt.Println(class.ID) //班级ID 3,4,5,6
	//}
	//
	//////指定批量添加的大小
	////db.CreateInBatches(classes,100)
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"msg": "添加成功!",
	//	"cl":  classes,
	//})

	////多对多的添加
	////实例化学生对象
	//var stu Student
	//stu = Student{BaseModel: BaseModel{Name: "dianDian"}, Sno: 200111, Pwd: "123", Tel: "111", Gender: 1, ClassID: 3}
	////添加前的stu:  {{0 dianDian <nil> <nil> <nil>} 200111 123 111 1 <nil>  3 {{0  <nil> <nil> <nil>} 0 0 {{0  <nil> <nil> <nil>} 0   0 <nil> }} []}
	//fmt.Println("添加前的stu: ", stu)
	//db.Create(&stu)
	///*
	//	添加后的stu:  {{16 dianDian 2023-02-21 14:17:03.799 +0800 CST 2023-02-21 14:17:03.799 +0800 CST 2023-02-21 14:17:03.799 +0800 CST} 200111 123 111 1 <nil>  3 {{0  <nil> <nil> <nil>} 0 0 {{0  <nil> <nil> <nil>} 0   0 <nil> }} []}
	//	可以看到Course是空的
	//*/
	//fmt.Println("添加后的stu: ", stu)

	////多对多的绑定(学生和课程)
	////为学生绑定课程
	//var courses []Course
	////先查询
	//db.Where("id in ?", []int{1, 2}).Find(&courses)
	////打印查询出来的课程
	//fmt.Println("courses>>> ", courses)
	//
	////第一种绑定方法  创建学生的同时绑定课程
	//stu1 := Student{
	//	BaseModel: BaseModel{Name: "翟辰瑜"},
	//	Sno:       200112,
	//	Pwd:       "123",
	//	Tel:       "15556835561",
	//	Gender:    1,
	//	ClassID:   3,
	//	//绑定查到的课程
	//	Courses: courses,
	//}
	////插入到数据库中
	//db.Create(&stu1)

	////第二种  创建学生后 在绑定课程
	////为查理绑定课程
	//var courses []Course
	////先查询
	//db.Where("name in ?", []string{"数据结构导论", "数据库及其应用"}).Find(&courses)
	////打印查询出来的课程
	//fmt.Println("courses>>> ", courses)
	////查找学生
	//var stu Student
	////数据库查询
	//db.Where("name = ?", "查理").Take(&stu)
	////给查理绑定课程
	////db.Model查找修改当前学生stu.关联Courses找到第三张表student2courses这张表.追加要绑定的课程courses
	////db.Model(&stu).Association("Courses").Append(courses)

	////clear 取消关联,清楚关联的记录
	//db.Model(&stu).Association("Courses").Clear()

	////查询 关联关系
	////定义课程
	//var courses []Course
	////定义学生并查询
	//var stu Student
	//db.Where("name = ?", "翟辰瑜").Take(&stu)
	////查询
	//db.Model(&stu).Association("Courses").Find(&courses)
	//
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"stu":     stu,
	//	"courses": courses,
	//})

	//// 批量添加课程
	////实例化课程对象
	//course1 := Course{BaseModel: BaseModel{Name: "数字电路"}, Credit: 3, Period: 16, TeacherID: 8}
	//course2 := Course{
	//	BaseModel: BaseModel{Name: "计算机微型接口"},
	//	Credit:    2,
	//	Period:    18,
	//	TeacherID: 10,
	//}
	//course3 := Course{BaseModel: BaseModel{Name: "计算机网络"}, Credit: 3, Period: 17, TeacherID: 11}
	////创建一个课程切片
	//courses := []Course{course1, course2, course3}
	//
	////写入数据库表记录
	//db.Create(&courses)
	//
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"courses": courses,
	//})

	//添加多对多的关联表的记录
	//绑定课程对象切片
	var courses []Course
	//先查询
	//获取课程,写入课程对象切片
	db.Where("name in ?", []string{"数字电路", "计算机网络"}).Find(&courses)
	fmt.Println("courses>>> ", courses)

	////添加学生1
	////创建学生不绑定课程
	//stu1 := Student{BaseModel: BaseModel{Name: "幸尧声"}, Sno: 200121, Pwd: "123", Tel: "13858413483", Gender: 1, ClassID: 6}
	//db.Create(&stu1)
	//
	////多对多添加方式1
	////创建学生的同时绑定课程
	//stu2 := Student{
	//	BaseModel: BaseModel{Name: "颜淮锨"},
	//	Sno:       200122,
	//	Pwd:       "123",
	//	Tel:       "13792715256",
	//	Gender:    0,
	//	ClassID:   5,
	//	//给学生绑定课程
	//	Courses: courses,
	//}
	////添加到数据库
	//db.Create(&stu2)

	////多对多添加方式2
	////先创建学生在绑定课程
	//stu3 := Student{
	//	BaseModel: BaseModel{Name: "龙栋诚"},
	//	Sno:       200123,
	//	Pwd:       "123",
	//	Tel:       "13896458484",
	//	Gender:    1,
	//	ClassID:   6,
	//}
	////添加学生到数据库
	//db.Create(&stu3)
	////打印学生ID
	//fmt.Println("stu3的ID>>> ", stu3.ID)
	////给学生关联绑定课程
	////注意: Courses是多对多的关联字段,不是关联表
	//db.Model(&stu3).Association("Courses").Append(courses)

	//多对多添加方式3
	//已有学生关联绑定课程
	//先查询学生
	var stu4 Student
	db.Where("name = ?", "羿伊琳").First(&stu4)
	//打印查询到的学生信息
	fmt.Println("stu4 的ID >>> ", stu4.ID)
	//关联绑定课程
	db.Model(&stu4).Association("Courses").Append(courses)

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

}

//Group Having 分组查询结构体

type Result struct {
	ClassID int
	Total   int
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

	////定义学生变量
	//var student []Student
	////Preload预加载Class这张表 ,预加载的表是跟后边查询的表必须是相关的
	//db.Preload("Class").Find(&student)
	//fmt.Println("students>>> ", student)
	//
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"student": student,
	//})

	//基于string的where语句
	//var student Student
	////查询名字叫做查理的学生
	//// SELECT * FROM `students` WHERE name = '查理' ORDER BY `students`.`id` LIMIT 1
	//db.Where("name = ?", "查理").First(&student)
	//fmt.Println("学生姓名>>> ", student.Name)

	////查询学号在2001到200105之间的学生
	//var student []Student
	//SELECT * FROM `students` WHERE Sno between 200101 and 200105
	//db.Where("Sno between ? and ?", 200101, 200105).Find(&student)
	//fmt.Println(student)

	////查询学号在200105,200107的学生
	//var student []Student
	////SELECT * FROM `students` WHERE Sno in (200105,200107)
	//db.Where("Sno in ?", []int64{200105, 200107}).Find(&student)
	//fmt.Println(student)

	////模糊查询
	//var student []Student
	////查询姓氏为刘姓的学生
	////SELECT * FROM `students` WHERE name like '刘%'
	//db.Where("name like ?", "刘%").Find(&student)
	//fmt.Println(student)

	////查询创建时间大于 2023-03-01 12:29:00  的学生
	//var student []Student
	////SELECT * FROM `students` WHERE create_time > '2023-03-01 12:29:00'
	//db.Where("create_time > ?", "2023-03-01 12:29:00").Find(&student)
	//fmt.Println(student)

	////Struct条件
	//var student []Student
	////结构体查询
	////SELECT * FROM `students` WHERE `students`.`name` = '查理' AND `students`.`gender` = 1
	//db.Where(&Student{BaseModel: BaseModel{Name: "查理"}, Gender: 1}).Find(&student)
	////注意:使用结构体作为条件查询时,GORM只会查询非零值字段
	//// SELECT * FROM `students` WHERE `students`.`name` = '羿伊琳'  零值不会作为字段查询
	//db.Where(&Student{BaseModel: BaseModel{Name: "羿伊琳"}, Gender: 0}).Find(&student)
	//fmt.Println(student)

	////map条件
	//var student []Student
	////SELECT * FROM `students` WHERE `Gender` = 1 AND `Name` = '查理'
	//db.Where(map[string]interface{}{"Name": "查理", "Gender": 1}).Find(&student)
	//fmt.Println(student)
	////使用map条件非零值字段也会查询
	////SELECT * FROM `students` WHERE `Gender` = 0 AND `Name` = '羿伊琳'
	//db.Where(map[string]interface{}{"Name": "羿伊琳", "Gender": 0}).Find(&student)
	//fmt.Println(student)
	//
	////响应
	//ctx.JSON(http.StatusOK, gin.H{
	//	"stu": student,
	//})

	//其他查询语句
	var student []Student
	////select语句,表示选择,其中写SQL部分
	////SELECT name,sno FROM `students` WHERE id = 11 LIMIT 1
	//db.Select("name,sno").Where("id = ?", 11).Take(&student)
	//log.Println(student)
	//
	//db.Omit("name", "sno").Find(&student)
	//fmt.Println(student)

	////not语句 用法类似于where
	////查询学号不在200104到200123的学生
	//db.Not("sno between ? and ?", 200104, 200123).Find(&student)
	//fmt.Println(student)

	////or语句
	////查询学号为200102或者姓为庞姓的学生
	//db.Where("sno  = ?", 200102).Or("name like ?", "庞%").Find(&student)
	//fmt.Println(student)

	////Order表示排序,其中写SQL部分
	////SELECT * FROM `students` WHERE create_time >= '2023-02-20' ORDER BY create_time desc,id
	////创建时间大于等于2023-02-20 降序排列
	//db.Where("create_time >= ?", "2023-02-20").Order("create_time desc,id").Find(&student)
	//log.Println(student)

	////Limit Offset 分页常用
	////SELECT * FROM `students` ORDER BY create_time desc LIMIT 3 OFFSET 1
	//db.Order("create_time desc").Limit(3).Offset(1).Find(&student)
	//log.Println(student)

	////Count计算行数
	////SELECT count(*) FROM `students`
	//var total int64
	//db.Model(Student{}).Count(&total)
	//fmt.Println(total)

	//Group Having 分组查询,其中写SQL部分,Group必须和Select一起连用
	//创建一个result的对象切片results
	var results []Result
	//SELECT class_id,Count(*) as total FROM `students` GROUP BY `class_id` HAVING total > 0
	db.Model(Student{}).Select("class_id,Count(*) as total").Group("class_id").Having("total > 0").Scan(&results)

	log.Println(results)

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"stu": student,
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

func updateNew(ctx *gin.Context) {
	//Save更新某条记录的所有字段
	//实例化学生对象
	//stu01 := Student{}
	////先查
	//db.First(&stu01)
	//fmt.Println(stu01.Name)
	////修改查出来的学生的姓名:俞珑椒更改为俞小小
	//stu01.Name = "俞小小"
	//db.Save(&stu01)
	//fmt.Println(stu01.Name)
	///*不推荐使用,因为会更新所有字段
	//[1.099ms] [rows:1] SELECT * FROM `students` ORDER BY `students`.`id` LIMIT 1
	//俞珑椒
	//
	//2023/03/02 16:30:29 /Users/denis/code/golang_study/src/gocode/go10/day16_03 Gorm总结/07 数据库表记录的增删改查/main.go:718
	//[10.169ms] [rows:1] UPDATE `students` SET `name`='俞小小',`create_time`='2023-02-16 13:16:38.374',`update_time`='2023-02-16 13:16:38.374',`delete_time`='2023-02-16 13:16:38.374',`sno`=200103,`pwd`='123',`tel`='18802628595',`gender`=0,`birth`=NULL,`remark`='新生',`class_id`=4 WHERE `id` = 4
	//俞小小
	//*/

	////Update基于主键更新某条记录的单个字段
	////更新ID为4学生俞小小的名字为俞珑椒
	//stu02 := Student{BaseModel: BaseModel{ID: 4}}
	////数据库更新
	////UPDATE `students` SET `name`='俞珑椒' WHERE `id` = 4
	//db.Model(&stu02).Update("name", "俞珑椒")
	////打印ID为4的学生姓名
	//fmt.Println("学生姓名: ", stu02.Name)

	////Update更新所有记录的单个字段
	////更改所有的课程周期为20周
	////UPDATE `courses` SET `period`='20'
	//db.Model(&Course{}).Update("period", 20)

	////Update自定义条件而非主键记录更新某字段
	////UPDATE `students` SET `sno`='200102' WHERE name = '俞珑椒'
	//db.Model(&Student{}).Where("name = ? ", "俞珑椒").Update("sno", "200102")

	////Update 更新多个字段
	////通过`struct`更新多个字段,不会更新零值字段
	////更新学生ID为4的sno为200103 和 pwd为321
	////UPDATE `students` SET `sno`=200103,`pwd`='321' WHERE id = 4
	//db.Model(&Student{}).Where("id = ?", 4).Updates(Student{Sno: 200103, Pwd: "321"})
	//通过`map`更新多个字段,零值字段也会更新
	//更新学生ID为9学生的学号为200102和remark为新生
	//UPDATE `students` SET `remark`='新生',`sno`=200102 WHERE id = 9
	//db.Model(&Student{}).Where("id = ? ", 9).Updates(map[string]interface{}{"Sno": 200102, "Remark": "新生"})

	//更新表达式
	//更新班级人数加1
	//UPDATE `classes` SET `num`=Num + 1 WHERE id = 3
	db.Model(&Class{}).Where("id = ?", 3).Update("Num", gorm.Expr("Num + 1"))

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

func getQuery(ctx *gin.Context) {

	////查询学生庞川灵的班级名称
	////方式一:手动查询
	////先查询学生
	////实例化学生对象
	//stu := Student{}
	////数据库查询
	////SELECT * FROM `students` WHERE name = '庞川灵' LIMIT 1
	//db.Where("name = ?", "庞川灵").Take(&stu)
	////打印学生姓名和班级ID
	//fmt.Println("学生姓名: ", stu.Name+"  班级ID: ", stu.ClassID)
	////通过classID查询班级
	////实例化班级对象
	//class := Class{}
	////数据库查询
	//// SELECT * FROM `classes` WHERE id = 5  LIMIT 1
	//db.Where("id = ? ", stu.ClassID).Take(&class)
	////打印班级名称
	//fmt.Println("班级名称: ", class.Name)

	////Preload预加载查询
	////实例化学生对象
	//stu := Student{}
	////数据库查询,查询的时候预加载class
	////和上边手动查询一样也是两条SQL
	////SELECT * FROM `classes` WHERE `classes`.`id` = 5
	////SELECT * FROM `students` WHERE name = '庞川灵' LIMIT 1
	//db.Where("name = ?", "庞川灵").Preload("Class").Take(&stu)
	////打印班级名称
	////班级名称:  计算机科学与技术一班
	//fmt.Println("班级名称: ", stu.Class.Name)

	////查询学生庞川灵所在班级名称和所报的课程
	////实例化学生对象
	//stu := Student{}
	////数据库查询
	///*
	//	SELECT * FROM `classes` WHERE `classes`.`id` = 5
	//	SELECT * FROM `student2course` WHERE `student2course`.`student_id` = 9
	//	SELECT * FROM `courses` WHERE `courses`.`id` = 1
	//	SELECT * FROM `students` WHERE name = '庞川灵' LIMIT 1
	//	班级名称:  计算机科学与技术一班
	//	课程名:  [{{1 数据结构导论 2023-02-18 21:13:18.074 +0800 CST 2023-02-18 21:13:18.074 +0800 CST 2023-02-18 21:13:18.074 +0800 CST} 3 20 9 {{0  <nil> <nil> <nil>} 0   0 <nil> }}]
	//*/
	//db.Where("name = ?", "庞川灵").Preload("Class").Preload("Courses").Find(&stu)
	////打印学生所在班级名称
	//fmt.Println("班级名称: ", stu.Class.Name)
	////打印学生所报课程
	//fmt.Println("课程信息: ", stu.Courses)
	//fmt.Println("第一个课程: ", stu.Courses[:1])
	////或者
	//db.Where("name = ?", "庞川灵").Preload("Class.Courses").Find(&stu)
	////打印学生所在班级名称
	//fmt.Println("班级名称: ", stu.Class.Name)
	////打印学生所报课程
	//fmt.Println("课程信息: ", stu.Courses)
	//fmt.Println("第一个课程: ", stu.Courses[:1])
	//////有多个课程时,循环展示每个课程
	////stuCourses := stu.Courses
	////for k, v := range stuCourses {
	////	fmt.Println(k, v)
	////}

	////实例化班级对象
	//class := Class{}
	////数据库查询
	//db.Where("name = ?", "软件一班").Preload("Student").Find(&class)
	////打印class信息
	//fmt.Println("class>>> ", class)
	//fmt.Println("class_student>>> ", class.Student)
	////循环展示学生
	//for _, stu := range class.Student {
	//	fmt.Println("学生ID: ", stu.ID, "   学生姓名: "+stu.Name)
	//}

	////实例化班级对象
	//class := Class{}
	////数据库查询+预加载Preload
	//db.Where("name = ? ", "软件一班").Preload("Student.Courses").Find(&class)
	////打印班级数据
	//fmt.Println("class>>> ", class)
	//fmt.Println("class_student>>> ", class.Student)
	////循环
	//for _, stu := range class.Student {
	//	fmt.Println("学生ID>>> ", stu.ID)
	//	fmt.Println("学生姓名>>> ", stu.Name)
	//	fmt.Println("学生所报课程信息>>> ", stu.Courses)
	//	fmt.Println("-----------------------------------------------------------")
	//	for _, course := range stu.Courses {
	//		fmt.Println(course)
	//	}
	//}

	////查询那些学生报了数据结构课程
	////实例化课程对象
	//course := Course{}
	////数据库查询
	//db.Where("name = ? ", "数据结构导论").Preload("Students").Find(&course)
	////打印课程信息
	//fmt.Println("course>>>", course)
	//fmt.Println("course_students>>>", course.Students)
	////循环
	//for _, stu := range course.Students {
	//	fmt.Println("报了数据结构课程的学生ID>>> ", stu.ID)
	//	fmt.Println("报了数据结构课程的学生姓名>>> ", stu.Name)
	//
	//}

	////查询学生庞川灵的班级名称(一对多)
	////实例化学生对象
	//var stu = Student{}
	////数据库查询
	//db.Where("students.name = ?", "庞川灵").Joins("Class").Find(&stu)
	////打印学生信息
	//fmt.Println("stu>>> ", stu)
	//fmt.Println("stu_class>>> ", stu.Class)
	//fmt.Println("stu_class_name>>> ", stu.Class.Name)
	//
	///*
	//	[1.311ms] [rows:1] SELECT `students`.`id`,`students`.`name`,`students`.`create_time`,`students`.`update_time`,`students`.`delete_time`,`students`.`sno`,`students`.`pwd`,`students`.`tel`,`students`.`gender`,`students`.`birth`,`students`.`remark`,`students`.`class_id`,`Class`.`id` AS `Class__id`,`Class`.`name` AS `Class__name`,`Class`.`create_time` AS `Class__create_time`,`Class`.`update_time` AS `Class__update_time`,`Class`.`delete_time` AS `Class__delete_time`,`Class`.`num` AS `Class__num`,`Class`.`tutor_id` AS `Class__tutor_id` FROM `students` LEFT JOIN `classes` `Class` ON `students`.`class_id` = `Class`.`id` WHERE students.name = '庞川灵'
	//	stu>>>  {{9 庞川灵 2023-02-17 22:58:58.801 +0800 CST 2023-02-17 22:58:58.801 +0800 CST 2023-02-17 22:58:58.801 +0800 CST} 200102 123 18600647479 1 <nil> 新生 5 {{5 计算机科学与技术一班 2023-02-15 10:40:26.61 +0800 CST 2023-02-15 10:40:26.61 +0800 CST 2023-02-15 10:40:26.61 +0800 CST} 39 3 {{0  <nil> <nil> <nil>} 0   0 <nil> } []} []}
	//	stu_class>>>  {{5 计算机科学与技术一班 2023-02-15 10:40:26.61 +0800 CST 2023-02-15 10:40:26.61 +0800 CST 2023-02-15 10:40:26.61 +0800 CST} 39 3 {{0  <nil> <nil> <nil>} 0   0 <nil> } []}
	//	stu_class_name>>>  计算机科学与技术一班
	//*/

	//查询ID为13学生所选的课程(多对多)
	//实例化学生对象
	stu := Student{BaseModel: BaseModel{ID: 13}}
	//定义课程对象
	var course []Course //一个学生会有多个课程,所以使用切片
	//数据库查询
	//注意db.Model()中必须是ID筛选
	/*
		SELECT `courses`.`id`,`courses`.`name`,`courses`.`create_time`,`courses`.`update_time`,`courses`.`delete_time`,`courses`.`credit`,`courses`.`period`,`courses`.`teacher_id` FROM `courses` JOIN `student2course` ON `student2course`.`course_id` = `courses`.`id` AND `student2course`.`student_id` = 13
		courses>>>  [{{1 数据结构导论 2023-02-18 21:13:18.074 +0800 CST 2023-02-18 21:13:18.074 +0800 CST 2023-02-18 21:13:18.074 +0800 CST} 3 20 9 {{0  <nil> <nil> <nil>} 0   0 <nil> } []} {{5 数字电路 2023-03-01 11:47:27.242 +0800 CST 2023-03-01 11:47:27.242 +0800 CST 2023-03-01 11:47:27.242 +0800 CST} 3 16 8 {{0  <nil> <nil> <nil>} 0   0 <nil> } []} {{7 计算机网络 2023-03-01 11:47:27.242 +0800 CST 2023-03-01 11:47:27.242 +0800 CST 2023-03-01 11:47:27.242 +0800 CST} 3 17 11 {{0  <nil> <nil> <nil>} 0   0 <nil> } []}]
	*/
	db.Model(&stu).Association("Courses").Find(&course)
	//打印课程信息
	fmt.Println("courses>>> ", course)
	//循环
	for _, course := range course {
		fmt.Println("课程名>>> ", course.Name)
	}

	//响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "查询成功",
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

	r.GET("query", getQuery)

	//更新
	r.GET("/update", Update)

	//NEW update
	r.GET("/updateNew", updateNew)

	//删除
	r.GET("/delete", Delete)

	//启动
	r.Run()
}
