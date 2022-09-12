package main

import (
	"fmt"
	"strings"
)

//客户信息管理系统结构体版

//定义Sid
var Sid = 0

//客户信息管理数据结构体
type Cms struct {
	Cid    int
	Name   string
	Gender string
	Age    int8
	Email  string
}

var User []Cms

////自增ID函数
//func ConstId(id int) int {
//	index := -1
//	for i := 0; i < len(User); i++ {
//		if User[i].Cid == id {
//			index = i
//		}
//	}
//	return index1
//}

//返回上一层函数
func isBack() bool {
	var xuanZE string
	fmt.Println("请选择是否返回上一层[Y:N]: ")
	fmt.Scan(&xuanZE)
	if strings.ToUpper(xuanZE) == "Y" {
		return true
	} else {
		return false
	}
}

//客户输入信息函数
func inputUser() (name, gender string, age int8, email string) {
	fmt.Println("请输入客户姓名: ")
	fmt.Scan(&name)
	fmt.Println("请输入客户性别: ")
	fmt.Scan(&gender)
	fmt.Println("请输入客户年龄: ")
	fmt.Scan(&age)
	fmt.Println("请输入客户邮箱: ")
	fmt.Scan(&email)

	return name, gender, age, email
}

//添加客户信息函数
func userAdd() {
	for true {
		fmt.Println("-----------------添加客户信息开始-----------------")
		name, gender, age, email := inputUser()

		//自增ID
		Sid++

		NewUser := Cms{
			Cid:    Sid,
			Name:   name,
			Gender: gender,
			Age:    age,
			Email:  email,
		}
		User = append(User, NewUser)
		//fmt.Println(User)
		fmt.Println("-----------------添加客户信息结束-----------------")

		//是否返回上一层
		b := isBack()
		if b {
			break
		}
	}
}

func chaUser() {
	for true {
		fmt.Println("-----------------查看客户信息开始-----------------")
		//fmt.Println(User)
		for _, v := range User {
			fmt.Printf("客户信息编号:%-8d 姓名:%-8s 性别:%-8s 年龄:%-8d 邮箱:%-8s\n", v.Cid, v.Name, v.Gender, v.Age, v.Email)
		}
		fmt.Println("-----------------查看客户信息结束-----------------")

		b := isBack()
		if b {
			break
		}
	}
}

func main() {
	for true {
		wel := `
-----------------客户信息管理系统(结构体版)-----------------
			1. 添加客户
			2. 查看客户
			3. 更新客户
			4. 删除客户
			5. 保存客户
			6. 查询已保存的客户信息
			7. 退出客户信息管理系统
------------------------------------------------------------
`
		fmt.Println(wel)
		var choice string
		fmt.Println("请输入您的选择 [1-7]: ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			{
				//添加客户功能
				userAdd()
			}
		case "2":
			{
				//查看客户功能
				chaUser()
			}
		case "3":
			{
				//更新客户功能
				fmt.Println("3")
			}
		case "4":
			{
				//删除客户功能
				fmt.Println("4")
			}
		case "5":
			{
				//保存客户功能
				fmt.Println("5")
			}
		case "6":
			{
				//查看已保存的客户功能
				fmt.Println("6")
			}
		case "7":
			{
				//退出客户信息管理系统功能
				fmt.Println("您已退出客户信息管理系统(结构体版),欢迎再次使用!!!")
				break
			}
		default:
			//输入的选择数值不合法
			fmt.Println("输入的选择数值不合法,请再次输入...")
			continue
		}

	}

}
