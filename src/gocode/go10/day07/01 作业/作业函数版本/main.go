package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// 客户信息管理系统欢迎界面函数
func wel() {
	wel := `
------------------客户信息管理系统------------------
			1. 添加客户
			2. 查看客户
			3. 更新客户
			4. 删除客户
			5. 保存客户
			6. 退出
------------------------------------------------------
`
	fmt.Println(wel)
}

// 客户信息管理系统选择功能函数
func choice() {
	for true {
		var choice string
		wel()
		fmt.Println("请输入您的选择[1-5]: ")
		fmt.Scan(&choice)

		//选择分支判断Switch
		switch choice {
		case "1":
			{
				//添加客户功能
				userAdd()
			}
		case "2":
			{
				//查看客户功能
				queryUser()
			}
		case "3":
			{
				//更新客户功能
				updateUser()
			}
		case "4":
			{
				//删除客户功能
				delUser()
			}
		case "5":
			{
				//保存客户功能
				saveUser()
			}
		case "6":
			{
				//退出客户信息管理系统
				fmt.Println("您已退出客户信息管理系统,欢迎再次使用...")
				break
			}
		default:
			fmt.Println("您输入的选择数值非法,请您重新输入,,,")
			continue
		}
	}

}

// 返回上一层功能函数
func isBack() bool {
	var back string
	fmt.Println("是否选择返回上一层[Y/N]: ")
	fmt.Scan(&back)
	if strings.ToUpper(back) == "Y" {
		return true
	} else {
		return false
	}
}

// 客户信息管理系统数据结构
var user []map[string]interface{}

// 定义自增sId
var sId int

// 自增id切片索引功能函数
func custId(id int) int {
	var index = -1
	for i := 0; i < len(user); i++ {
		if user[i]["cid"] == id {
			index = i
		}
	}
	return index
}

// 用户输入函数
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

// 添加用户功能函数
func userAdd() {
	for true {
		fmt.Println("------------------添加客户信息开始------------------")
		name, gender, age, email := inputUser()

		//sId自加1
		sId++

		newUser := map[string]interface{}{
			"cid":    sId,
			"name":   name,
			"gender": gender,
			"age":    age,
			"email":  email,
		}
		user = append(user, newUser)
		//fmt.Println(user)  //测试添加客户功能是否正常
		fmt.Println("------------------添加客户信息结束------------------")

		b := isBack()
		if b {
			break
		}

	}

}

// 查询客户功能函数
func queryUser() {
	for true {
		fmt.Println("------------------查询客户信息开始------------------")
		for _, info := range user {
			//fmt.Println(info)
			fmt.Printf("客户编号:%-8d 姓名:%-8s 性别:%-8s 年龄:%-8d 邮箱:%-8s\n", info["cid"], info["name"], info["gender"], info["age"], info["email"])
		}
		fmt.Println("------------------查询客户信息结束------------------")

		b := isBack()
		if b {
			break
		}
	}

}

// 更新客户功能函数
func updateUser() {
	for true {
		fmt.Println("------------------更新客户信息开始------------------")
		var updataId int
		fmt.Println("请输入要更新修改的客户编号: ")
		fmt.Scan(&updataId)

		updataIndex := custId(updataId)
		//判断用户输入的客户编号是否为-1
		if updataIndex == -1 {
			fmt.Println("您输入的客户信息编号不存在,请重新输入xxx")
			continue
		}
		//修改客户信息
		fmt.Println("请输入要修改的客户信息: ")
		name, gender, age, email := inputUser()
		user[updataIndex]["name"] = name
		user[updataIndex]["gender"] = gender
		user[updataIndex]["age"] = age
		user[updataIndex]["email"] = email
		fmt.Println("------------------更新客户信息结束------------------")
		b := isBack()
		if b {
			break
		}
	}

}

// 删除客户功能函数
func delUser() {
	for true {
		fmt.Println("------------------删除客户信息开始------------------")
		var delId int
		fmt.Println("请输入要删除的客户编号: ")
		fmt.Scan(&delId)
		delIndex := custId(delId)
		if delIndex == -1 {
			fmt.Println("您输入的客户信息编号不存在,请重新输入xxx")
			continue
		}
		user = append(user[:delIndex], user[delIndex+1:]...)
		fmt.Println("------------------删除客户信息结束------------------")

		b := isBack()
		if b {
			break
		}
	}

}

// 保存客户信息功能函数
func saveUser() {
	for true {
		fmt.Println("------------------保存客户信息开始------------------")
		userData, err := json.Marshal(user)
		ioutil.WriteFile("userData.json", userData, 0666)
		if err != nil {
			fmt.Println("客户信息保存失败xxx")
		} else {
			fmt.Println("客户信息保存成功***")
		}
		fmt.Println("------------------保存客户信息结束------------------")

		b := isBack()
		if b {
			break
		}
	}

}

func main() {

	//调用客户信息管理系统选择函数
	choice()

}
