create database css default character set utf8mb4;

show databases;

# 结构体 映射 table表

# create table t1(
#     name varchar(255) unique not null ,
#     age int
# )

# type t1 struct {
# 	Name string `gorm:"type:varchar(32);unique;not null"`
# 	Age  int
# }

#学生选课系统

#老师表:姓名,年龄,手机号,创建时间,更新时间,删除时间

#班级表:班级名,班级人数,导员,创建时间,更新时间,删除时间

#学生表:学号,姓名,年龄,性别,班级id,创建时间,更新时间,删除时间

#课程表:课程名称,学分,教师,周期,老师,创建时间,更新时间,删除时间

#关系表:主键,学生ID,课程ID,创建时间,更新时间,删除时间