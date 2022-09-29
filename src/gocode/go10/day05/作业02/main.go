package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

//构建客户信息管理系统数据结构
var sId = 0
var user []map[string]interface{}

//自增ID的函数
func custID(id int) int {
	index := -1
	for i := 0; i < len(user); i++ {
		if user[i]["cid"] == id {
			index = i
		}
	}
	return index
}

//返回上一层函数
func isBack() bool {
	var xuanZe string
	fmt.Println("请选择是否返回上一层 [Y/N]: ")
	fmt.Scan(&xuanZe)
	if strings.ToUpper(xuanZe) == "Y" {
		return true
	} else {
		return false
	}
}

//输入客户信息函数¬Œ
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

//添加客户信息函数
func userAdd() {
	for true {
		fmt.Println("-----------------添加客户信息开始-----------------")
		name, gender, age, email := inputUser()
		sId++
		newUser := map[string]interface{}{
			"cid":    sId,
			"name":   name,
			"gender": gender,
			"age":    age,
			"email":  email,
		}
		user = append(user, newUser)
		fmt.Println("-----------------添加客户信息结束-----------------")
		b := isBack()
		if b {
			break
		}

	}

}

//查看客户信息函数
func chaUser() {
	for true {
		fmt.Println("-----------------查看客户信息开始-----------------")
		for _, user := range user {
			fmt.Printf("客户信息编号:%-8d 姓名:%-8s 性别:%-8s 年龄:%-8s 邮箱:%-8s\n", user["cid"], user["name"], user["gender"], user["age"], user["email"])
		}
		fmt.Println("-----------------查看客户信息结束-----------------")
		b := isBack()
		if b {
			break
		}
	}

}

//更新客户信息函数
func updateUser() {
	for true {
		fmt.Println("-----------------修改客户信息开始-----------------")
		fmt.Println("请输入要修改的客户编号: ")
		var modID int
		fmt.Scan(&modID)
		modIndex := custID(modID)
		if modIndex == -1 {
			fmt.Println("您输入的客户信息编号不存在XXX")
			continue
		}
		fmt.Println("请输入要修改的客户信息: ")
		name, gender, age, email := inputUser()
		user[modIndex]["name"] = name
		user[modIndex]["gender"] = gender
		user[modIndex]["age"] = age
		user[modIndex]["email"] = email
		fmt.Println("-----------------修改客户信息结束-----------------")
		b := isBack()
		if b {
			break
		}
	}

}

//删除客户信息函数
func deleUser() {
	for true {
		var delteId int
		fmt.Println("-----------------删除客户信息开始-----------------")
		fmt.Println("请输出要删除的客户信息编号: ")
		fmt.Scan(&delteId)
		delteIndes := custID(delteId)
		if delteIndes == -1 {
			fmt.Println("您输入的客户信息编号不存在XXX")
			continue
		}
		user = append(user[:delteIndes], user[delteIndes+1:]...)
		fmt.Println("-----------------删除客户信息结束-----------------")
		b := isBack()
		if b {
			break
		}
	}

}

//保存客户信息函数
func cunUser() {
	for true {
		fmt.Println("-----------------保存客户信息开始-----------------")
		//fmt.Println(reflect.TypeOf(user))
		data, err := json.Marshal(user)
		ioutil.WriteFile("cmsData.json", data, 0666)
		if err != nil {
			fmt.Println("客户信息保存失败xxx")
		} else {
			fmt.Println("客户信息保存成功***")
		}
		fmt.Println("-----------------保存客户信息结束-----------------")
		b := isBack()
		if b {
			break
		}
	}

}

//查看已保存的客户信息函数
func chaYicun() {
	for true {
		fmt.Println("-----------------查看已保存客户信息开始-----------------")
		data1, err := ioutil.ReadFile("/Users/denis/code/golang_study/src/cmsData.json")
		if err != nil {
			fmt.Println("打开已保存客户信息失败xxx")
		} else {
			fmt.Println("打开已保存客户信息成功****")
		}
		//fmt.Println(reflect.TypeOf(data1)) //[]uint8
		var info = []map[string]interface{}{}
		json.Unmarshal(data1, &info)
		//fmt.Println(info)
		//[map[age:22 cid:1 email:dadadad@qq.com gender:male name:daerwen] map[age:12 cid:2 email:dadada@qq.com gender:male name:dafenqi]]
		for _, v := range info {
			fmt.Printf("客户信息编号:%#v  姓名:%-8s 性别:%-8s 年龄:%-8s 邮箱:%-8s\n", v["cid"], v["name"], v["gender"], v["age"], v["email"])
			//for i1, v1 := range v {
			//	fmt.Println(i1, v1)
			//}
		}
		//fmt.Printf("客户信息编号:%-8d 姓名:%-8s 性别:%-8s 年龄:%-8s 邮箱:%-8s\n",info[""])

		fmt.Println("-----------------查看已保存客户信息结束-----------------")
		b := isBack()
		if b {
			break
		}
	}

}
func main() {
	for true {

		var sele string
		wel := `
	------------------客户信息管理系统------------------
			1. 添加客户
			2. 查看客户
			3. 更新客户
			4. 删除客户
			5. 保存客户
			6. 查看已保存的客户信息
			7. 退出
	--------------------------------------------------
`
		//显示客户信息管理系统欢迎界面
		fmt.Println(wel)

		//引导用户输入
		fmt.Println("请输入您的选择[1-5]: ")
		fmt.Scan(&sele)

		//客户信息管理系统功能选择
		switch sele {
		case "1":
			{
				//添加客户信息
				userAdd()
			}
		case "2":
			{
				//查看客户信息
				chaUser()
			}
		case "3":
			{
				//修改客户信息
				updateUser()
			}
		case "4":
			{
				//删除客户信息
				deleUser()
			}
		case "5":
			{
				//保存客户信息
				cunUser()
			}
		case "6":
			{
				//查看已保存的客户信息
				chaYicun()
			}
		case "7":
			{
				//退出客户信息管理系统
				fmt.Println("您已退出客户信息管理系统,欢迎再次使用...")
				break
			}
		}
	}

}
