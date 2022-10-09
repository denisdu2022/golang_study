package model

import "fmt"

//定义客户信息管理系统数据结构

type Customer struct {
	Cid    int    `json:"cid"` //设置json别名,在存储到json文件时是小写的成员变量名
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
}

//客户信息输入函数
func InputCustomer() (name, gender string, age int, email string) {
	fmt.Print("请输入客户姓名: ")
	fmt.Scanln(&name)
	fmt.Print("请输入客户性别: ")
	fmt.Scanln(&gender)
	fmt.Print("请输入客户年龄: ")
	fmt.Scanln(&age)
	fmt.Print("请输入客户邮箱: ")
	fmt.Scanln(&email)
	return name, gender, age, email
}
