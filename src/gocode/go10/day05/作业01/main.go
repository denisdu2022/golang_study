package main

import (
	"fmt"
	"strings"
)

var sID int
var custNum []map[string]interface{}

func cmsId(id int) int {
	index := -1

	//遍历custnum切片
	for i := 0; i < len(custNum); i++ {
		if custNum[i]["cid"] == id {
			index = i
		}
	}
	return index
}

//返回上一层的函数
func isBack() bool {
	fmt.Println("请选择是否返回上一层[Y/N]: ")
	var backCh string
	fmt.Scan(&backCh)
	if strings.ToUpper(backCh) == "Y" {
		return true
	} else {
		return false
	}
}

//用户输入的函数
func inputUser() (name, gender, age, email string) {

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

//添加用户的函数
func addUser() {
	for true {
		fmt.Println("---------------添加客户开始---------------")
		name, gender, age, email := inputUser()

		//创建用户的map对象
		sID++ //用户自增ID
		newCustNum := map[string]interface{}{
			"cid":    sID,
			"name":   name,
			"gender": gender,
			"age":    age,
			"email":  email,
		}
		//添加用户的map对象添加到用户切片中
		custNum = append(custNum, newCustNum)
		b := isBack()
		if b {
			break
		}
		fmt.Println("---------------添加客户完成---------------")
	}
}

//查询用户信息的函数
func chaUser() {
	for true {
		fmt.Println("---------------客户列表开始---------------")
		for _, custNum := range custNum {
			fmt.Printf("编号: %-8d 姓名:%-8s 性别:%-8s 年龄:%-8s 邮箱:%-8s\n",
				custNum["cid"], custNum["name"], custNum["gender"], custNum["age"], custNum["email"])

		}
		fmt.Println("---------------客户列表结束---------------")
		b := isBack()
		if b {
			break
		}
	}

}

//修改用户信息的函数
func modUser() {
	for true {
		fmt.Println("---------------修改客户信息开始---------------")
		var modID int
		fmt.Println("请输入要修改的客户的编号: ")
		fmt.Scan(&modID)
		modIndex := cmsId(modID)
		if modIndex == -1 {
			fmt.Println("修改客户信息失败,客户编号ID不存在...")
			continue
		}
		fmt.Println("请输入要修改的客户信息: ")
		name, gender, age, emmail := inputUser()
		custNum[modIndex]["name"] = name
		custNum[modIndex]["gender"] = gender
		custNum[modIndex]["age"] = age
		custNum[modIndex]["email"] = emmail
		fmt.Println("---------------修改客户信息结束---------------")
		b := isBack()
		if b {
			break
		}
	}

}

//删除用户信息的函数
func deleteUser() {
	for true {
		fmt.Println("---------------删除客户信息开始---------------")
		var delId int
		fmt.Println("请输入要删除客户的编号: ")
		fmt.Scan(&delId)
		delIndex := cmsId(delId)
		if delIndex == -1 {
			fmt.Println("删除客户信息失败,客户编号ID不存在...")
			continue
		}
		custNum = append(custNum[:delIndex], custNum[delIndex+1:]...)
		fmt.Println("---------------删除客户信息结束---------------")
		b := isBack()
		if b {
			break
		}
	}

}

//存储用户信息的函数
func cunUser() {
	fmt.Println("次功能暂未上线,敬请期待......")
}

func main() {
	for true {
		var tiShi = `
----------------客户信息管理系统----------------
		1. 添加客户
		2. 查看客户
		3. 更新客户
		4. 删除客户
		5. 保存客户信息
		6. 退出
----------------------------------------------
`
		//显示提示信息
		fmt.Println(tiShi)

		//定义分支选择
		var xuanZe string
		fmt.Println("请输入您的选择: ")
		//输入分支选择
		fmt.Scan(&xuanZe)
		switch xuanZe {
		case "1":
			{
				addUser()
			}
		case "2":
			{
				chaUser()
			}
		case "3":
			{
				modUser()
			}
		case "4":
			{
				deleteUser()
			}
		case "5":
			{
				cunUser()
			}
		case "6":
			{
				break
			}
		}
		fmt.Println("您已退出客户管理系统...")
	}
}
