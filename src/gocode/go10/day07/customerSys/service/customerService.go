package service

import (
	"customerSys/model"
	"fmt"
	"os"
	"strings"
)

//客户信息管理系统数据结构Sys
type CustomersSys struct {
	Customer   []model.Customer
	CustomerID int
}

//返回上一层函数
func isBack() bool {
	var back string
	fmt.Println("请问是否返回上一层[Y/N]: ")
	fmt.Scan(&back)
	if strings.ToUpper(back) == "Y" {
		return true
	} else {
		return false
	}
}

//CustomerID函数
func (cs *CustomersSys) customerID(id int) int {
	indexCustomer := -1
	for i := 0; i < len(cs.Customer); i++ {
		for _, customerValue := range cs.Customer {
			if customerValue.Cid == id {
				indexCustomer = i
			}
		}
	}
	return indexCustomer
}

//客户信息管理系统欢迎界面
func Wel() {
	wel := `
-----------------------客户信息管理系统-----------------------
			1. 添加客户
			2. 查看客户
			3. 更新客户
			4. 删除客户
			5. 保存客户信息
			6. 退出客户信息管理系统
---------------------------------------------------------------------
`
	fmt.Println(wel)
}

//客户信息管理系统选择功能函数
func (cs *CustomersSys) Choice() {
	var Choice int
	fmt.Println("请输入您的选择[1-6]: ")
	fmt.Scan(&Choice)
	switch Choice {
	case 1:
		//添加客户功能
		cs.addCustomer()
	case 2:
		//查看客户功能
		cs.listCustomer()
	case 3:
		//更新客户功能
		cs.updateCustomer()
	case 4:
		//删除客户功能
		cs.delteCustomer()
	case 5:
		//保存客户信息功能
		cs.saceCustomer()
	case 6:
		//退出客户信息管理系统
		fmt.Println("您已退出客户信息管理系统,欢迎再次使用!!!")
		os.Exit(0)
	default:
		fmt.Println("您输入的选择数值非法,请重新输入...")

	}
}

//用户输入信息函数
func inputCustomer() (name, gender string, age int8, email string) {
	fmt.Println("请输入客户姓名:")
	fmt.Scan(&name)
	fmt.Println("请输入客户性别:")
	fmt.Scan(&gender)
	fmt.Println("请输入客户年龄:")
	fmt.Scan(&age)
	fmt.Println("请输入客户邮箱:")
	fmt.Scan(&email)
	return name, gender, age, email
}

//添加客户信息功能函数
func (cs *CustomersSys) addCustomer() {
	for true {
		fmt.Println("--------------------添加客户开始--------------------")
		//用户输入客户信息
		name, gender, age, email := inputCustomer()
		//每添加一个客户,CustomerID自加1
		cs.CustomerID++
		//定义一个新的输入的客户信息接收变量
		newCustomerInput := model.Customer{
			Cid:    cs.CustomerID,
			Name:   name,
			Gender: gender,
			Age:    age,
			Email:  email,
		}
		//将添加的用户追加到Customer
		cs.Customer = append(cs.Customer, newCustomerInput)
		fmt.Println("--------------------添加客户信息成功--------------------")
		//fmt.Println(cs.Customer)
		fmt.Println("--------------------添加客户开始--------------------")
		//判断用户是否返回上一层
		b := isBack()
		if b {
			break
		}
	}
}

//查询客户信息功能函数
func (cs *CustomersSys) listCustomer() {
	for true {
		fmt.Println("--------------------查询客户信息开始--------------------")
		fmt.Println("--------------------客户信息查询成功--------------------")
		for _, customerValues := range cs.Customer {
			fmt.Printf("客户编号:%-8d 姓名:%-8s 性别:%-8s 年龄:%-8d 邮箱:%-8s\n", customerValues.Cid, customerValues.Name,
				customerValues.Gender, customerValues.Age, customerValues.Email)
		}
		fmt.Println("--------------------查询客户信息结束--------------------")
		//判断用户是否返回上一层
		b := isBack()
		if b {
			break
		}
	}
}

//更新客户信息功能函数
func (cs *CustomersSys) updateCustomer() {
	for true {
		fmt.Println("--------------------更新客户信息开始--------------------")
		var cid int
		fmt.Println("请输入您要操作的客户信息编号: ")
		fmt.Scan(&cid)
		handleIndex := cs.customerID(cid)
		if handleIndex == -1 {
			fmt.Println("您输入的客户信息编号不存在,请重新输入...")
			continue
		}

		updateCustomer := &cs.Customer[handleIndex]
		var (
			name   string
			gender string
			age    int8
			email  string
		)
		fmt.Printf("请输入要操作的客户姓名(%s): ", updateCustomer.Name)
		fmt.Scanln(&name)
		fmt.Printf("请输入要操作的客户性别(%s): ", updateCustomer.Gender)
		fmt.Scanln(&gender)
		fmt.Printf("请输入要操作的客户年龄(%d): ", updateCustomer.Age)
		fmt.Scanln(&age)
		fmt.Printf("请输入要操作的客户邮箱(%s): ", updateCustomer.Email)
		fmt.Scanln(&email)

		if name != "" {
			updateCustomer.Name = name
		}
		if gender != "" {
			updateCustomer.Gender = gender
		}
		if age != 0 {
			updateCustomer.Age = age
		}
		if email != "" {
			updateCustomer.Email = email
		}
		fmt.Println("--------------------更新客户信息成功--------------------")
		fmt.Println("--------------------更新客户信息结束--------------------")
		//判断用户是否返回上一层
		b := isBack()
		if b {
			break
		}
	}

}

//删除客户信息功能函数
func (cs *CustomersSys) delteCustomer() {
	for true {
		fmt.Println("--------------------删除客户信息开始--------------------")
		var cid int
		fmt.Println("请输入您要操作的客户信息编号: ")
		fmt.Scan(&cid)
		handleIndex := cs.customerID(cid)
		if handleIndex == -1 {
			fmt.Println("您输入的客户信息编号不存在,请重新输入...")
			continue
		}
		cs.Customer = append(cs.Customer[:handleIndex], cs.Customer[handleIndex+1:]...)
		fmt.Println("--------------------删除客户信息成功--------------------")
		fmt.Println("--------------------删除客户信息结束--------------------")
		//判断用户是否返回上一层
		b := isBack()
		if b {
			break
		}
	}
}

//保存客户信息功能函数
func (cs *CustomersSys) saceCustomer() {
	for true {
		//判断用户是否返回上一层
		b := isBack()
		if b {
			break
		}
	}
}
