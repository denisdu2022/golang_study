package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 定义客户信息管理系统数据结构
type Customer struct {
	Cid    int
	Name   string
	Age    int8
	Gender string
	Email  string
}

//var customers []Customer

// 定义自增ID,新增一个用户的时候自加1
// var customerID int
// 定义客户信息管理Service
type CustomerService struct {
	customers  []Customer
	customerID int
}

func (cs *CustomerService) addCustomer() {
	for true {
		//添加客户功能
		//引导用户输入客户信息
		fmt.Printf("\033[1;35;40m%s\033[0m\n", "-----------------添加客户信息开始-----------------")

		var name, gender, email string
		var age int8
		fmt.Println("请输入客户姓名:")
		fmt.Scan(&name)
		fmt.Println("请输入客户性别:")
		fmt.Scan(&gender)
		fmt.Println("请输入客户年龄")
		fmt.Scan(&age)
		fmt.Println("请输入客户邮箱:")
		fmt.Scan(&email)

		//每新增一个客户,客户编号自加1
		cs.customerID++

		//定义新增的客户信息类型并赋值
		newCustomer := Customer{
			Cid:    cs.customerID,
			Name:   name,
			Age:    age,
			Gender: gender,
			Email:  email,
		}
		//将新增的客户信息追加到已经定义好的客户信息数据结构中
		cs.customers = append(cs.customers, newCustomer)

		//fmt.Println(customers)
		fmt.Printf("\033[1;35;40m%s\033[0m\n", "-----------------添加客户信息结束-----------------")
		fmt.Println()

		//判断用户是否继续输入,如果不继续输入则返回上一层
		var CustomerBack string
		fmt.Printf("\033[1;35;40m%s\033[0m\n", "请问是否返回上一层[Y/N]: ")
		fmt.Scan(&CustomerBack)
		if strings.ToUpper(CustomerBack) == "Y" {
			break
		} else {
			continue
		}

	}
}

func (cs CustomerService) listCustomer() {
	for true {
		//查看客户功能
		fmt.Println("-----------------查看客户信息开始-----------------")
		//for _, customerStructValue := range customers {
		//	fmt.Printf("编号:%-6d 姓名:%-6s 性别:%-6s 年龄:%-6d 邮箱:%-6s customers的地址: %p\n", customerStructValue.Cid, customerStructValue.Name,
		//		customerStructValue.Gender, customerStructValue.Age, customerStructValue.Email, &customers)
		//} //customers的地址: 0x102a26520
		for _, customerStructValue := range cs.customers {
			fmt.Printf("编号:%-6d 姓名:%-6s 性别:%-6s 年龄:%-6d 邮箱:%-6s\n", customerStructValue.Cid, customerStructValue.Name,
				customerStructValue.Gender, customerStructValue.Age, customerStructValue.Email)
		} //customers的地址: 0x102a26520
		fmt.Println("-----------------查看客户信息结束-----------------")
		fmt.Println()
		//判断用户是否继续输入,如果不继续输入则返回上一层
		var CustomerBack string
		fmt.Printf("\033[1;35;40m%s\033[0m\n", "请问是否返回上一层[Y/N]: ")
		fmt.Scan(&CustomerBack)
		if strings.ToUpper(CustomerBack) == "Y" {
			break
		} else {
			continue
		}
	}
}

func (cs CustomerService) updataCustomer() {
	for true {
		//更新客户功能
		fmt.Println("-----------------更新客户信息开始-----------------")
		//引导用户输入要修改的客户信息编号
		var cid int
		fmt.Println("请输入您要修改的客户信息编号: ")
		fmt.Scan(&cid)

		//定义updataIndex索引为-1 一般-1表示不存在
		var updataIndex = -1

		for indexCustomer, customerStructValue := range cs.customers {
			if customerStructValue.Cid == cid {
				updataIndex = indexCustomer
			}
		}
		//判断输入的客户信息编号是否存在
		if updataIndex == -1 {
			fmt.Println("您输入的客户信息编号不存在,请您重新输入: ")
			continue
		}

		//更新修改客户信息
		//将输入修改编号的的客户信息重新赋值给update
		//updataCustomer := customers[updataIndex]
		//fmt.Printf("updataCustomer的地址: %p\n", &updataCustomer) //updataCustomer的地址: 0x1400010c000
		//结构体是值类型,所以以上修改了值,并不会修改源customres的值,所以需要以下的写法
		//var updataCustomer *Customer
		//updataCustomer = &customers[updataIndex]

		//或者
		updataCustomer := &cs.customers[updataIndex]

		var name, gender, email string
		var age int8
		fmt.Printf("请输入要修改的客户姓名(%s):", updataCustomer.Name) //*(updataCustomer).Name 语法糖
		fmt.Scanln(&name)
		fmt.Printf("请输入要修改的客户性别(%s):", updataCustomer.Gender)
		fmt.Scanln(&gender)
		fmt.Printf("请输入要修改的客户年龄(%d):", updataCustomer.Age)
		fmt.Scanln(&age)
		fmt.Printf("请输入要修改的客户邮箱(%s):", updataCustomer.Email)
		fmt.Scanln(&email)
		fmt.Println("-----------------更新客户信息结束-----------------")

		//判断客户信息每一项是否都需要修改
		if name != "" {
			updataCustomer.Name = name
		}
		if gender != "" {
			updataCustomer.Gender = gender
		}
		if age != 0 {
			updataCustomer.Age = age
		}
		if email != "" {
			updataCustomer.Email = email
		}

		//判断用户是否继续输入,如果不继续输入则返回上一层
		var CustomerBack string
		fmt.Printf("\033[1;35;40m%s\033[0m\n", "请问是否返回上一层[Y/N]: ")
		fmt.Scan(&CustomerBack)
		if strings.ToUpper(CustomerBack) == "Y" {
			break
		} else {
			continue
		}
	}

}

func (cs *CustomerService) delteCustomer() {
	//删除客户功能
	for true {
		fmt.Println("-----------------删除客户信息开始-----------------")
		var delteCustomerID int
		fmt.Println("请输入要删除的客户信息的编号: ")
		fmt.Scan(&delteCustomerID)

		//定义updataIndex索引为-1 一般-1表示不存在
		var delteIndex = -1

		for indexCustomer, customerValue := range cs.customers {
			if customerValue.Cid == delteCustomerID {
				delteIndex = indexCustomer
			}
		}
		//判断输入的客户信息编号是否存在
		if delteIndex == -1 {
			fmt.Println("您输入的客户信息编号不存在,请您重新输入: ")
			continue
		}

		cs.customers = append(cs.customers[:delteIndex], cs.customers[delteIndex+1:]...)
		fmt.Println("-----------------删除客户信息结束-----------------")

		//判断用户是否继续输入,如果不继续输入则返回上一层
		var CustomerBack string
		fmt.Printf("\033[1;35;40m%s\033[0m\n", "请问是否返回上一层[Y/N]: ")
		fmt.Scan(&CustomerBack)
		if strings.ToUpper(CustomerBack) == "Y" {
			break
		} else {
			continue
		}
	}
}

func (cs CustomerService) saveCustomer() {
	//保存客户功能
	fmt.Println("------------------保存客户信息开始------------------")
	//json序列化客户信息数据
	jsonData, err := json.Marshal(cs.customers)
	ioutil.WriteFile("jsonCustomerdata.json", jsonData, 0666)
	if err != nil {
		fmt.Println("文件保存失败xxx")
	} else {
		fmt.Println("文件保存成功***")
	}
	fmt.Println(string(jsonData))
	fmt.Println("------------------保存客户信息结束------------------")
}

func main() {

	//实例化CustomerService
	cs := CustomerService{
		customers:  []Customer{},
		customerID: 0,
	}

	for true {
		fmt.Printf("\033[1;30;42m%s\033[0m\n", `
-----------------客户信息管理系统-----------------
			1. 添加客户
			2. 查看客户
			3. 更新客户
			4. 删除客户
			5. 保存客户信息
			6. 退出	
---------------------------------------------------
`)
		var choice string
		fmt.Println("请输入您的选择[1-6]: ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			{
				cs.addCustomer()
			}
		case "2":
			{
				cs.listCustomer()
			}
		case "3":
			{
				cs.updataCustomer()
			}
		case "4":
			{
				cs.delteCustomer()
			}
		case "5":
			{
				cs.saveCustomer()
			}
		case "6":
			{
				//退出客户信息管理系统
				fmt.Println("您已退出客户信息管理系统,欢迎再次使用!!!")
				os.Exit(0)
			}
		default:
			//非法输入
			fmt.Println("您输入的选择数值不合法,请您重新输入XXX")

		}
	}

}
