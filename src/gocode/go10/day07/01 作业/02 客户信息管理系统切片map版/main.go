package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//客户信息管理系统数据结构
var customers = []map[string]string{}

//定义自增ID
var customerID int

func main() {
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
		fmt.Println("请输入您的选择[1-7]: ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			{
				for true {
					//添加客户功能
					//引导用户输入客户信息
					fmt.Printf("\033[1;35;40m%s\033[0m\n", "-----------------添加客户信息开始-----------------")

					var name, gender, age, email string
					fmt.Println("请输入客户姓名:")
					fmt.Scan(&name)
					fmt.Println("请输入客户性别:")
					fmt.Scan(&gender)
					fmt.Println("请输入客户年龄")
					fmt.Scan(&age)
					fmt.Println("请输入客户邮箱:")
					fmt.Scan(&email)

					//每新增一个客户,客户编号自加1
					customerID++

					//定义新增的客户信息类型并赋值
					newCustomers := map[string]string{
						"cid":    strconv.Itoa(customerID),
						"name":   name,
						"gender": gender,
						"age":    age,
						"email":  email,
					}
					//将新增的客户信息追加到已经定义好的客户信息数据结构中
					customers = append(customers, newCustomers)
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
		case "2":
			{
				for true {
					//查看客户功能
					fmt.Println("-----------------查看客户信息开始-----------------")
					for _, customerValue := range customers {
						fmt.Printf("编号:%-6s 姓名:%-6s 性别:%-6s 年龄:%-6s 邮箱:%-6s\n", customerValue["cid"], customerValue["name"],
							customerValue["gender"], customerValue["age"], customerValue["email"])
					}
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
		case "3":
			{
				for true {
					//更新客户功能
					fmt.Println("-----------------更新客户信息开始-----------------")
					//引导用户输入要修改的客户信息编号
					var cid string
					fmt.Println("请输入您要修改的客户信息编号: ")
					fmt.Scan(&cid)

					//定义updataIndex索引为-1 一般-1表示不存在
					var updataIndex = -1

					for indexCustomer, customerValue := range customers {
						if customerValue["cid"] == cid {
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
					updataCustomer := customers[updataIndex]
					//var name, gender, age, email string
					fmt.Printf("请输入要修改的客户姓名(%s):", updataCustomer["name"])
					var name string
					fmt.Scanln(&name)
					fmt.Printf("请输入要修改的客户性别(%s):", updataCustomer["gender"])
					var gender string
					fmt.Scanln(&gender)
					fmt.Printf("请输入要修改的客户年龄(%s):", updataCustomer["age"])
					var age string
					fmt.Scanln(&age)
					fmt.Printf("请输入要修改的客户邮箱(%s):", updataCustomer["email"])
					var email string
					fmt.Scanln(&email)
					fmt.Println("-----------------更新客户信息结束-----------------")

					//判断客户信息每一项是否都需要修改
					if name != "" {
						updataCustomer["name"] = name
					}
					if gender != "" {
						updataCustomer["gender"] = gender
					}
					if age != "" {
						updataCustomer["age"] = age
					}
					if email != "" {
						updataCustomer["email"] = email
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
		case "4":
			{
				//删除客户功能
				for true {
					fmt.Println("-----------------删除客户信息开始-----------------")
					var delteCustomerID string
					fmt.Println("请输入要删除的客户信息的编号: ")
					fmt.Scan(&delteCustomerID)

					//定义updataIndex索引为-1 一般-1表示不存在
					var delteIndex = -1

					for indexCustomer, customerValue := range customers {
						if customerValue["cid"] == delteCustomerID {
							delteIndex = indexCustomer
						}
					}
					//判断输入的客户信息编号是否存在
					if delteIndex == -1 {
						fmt.Println("您输入的客户信息编号不存在,请您重新输入: ")
						continue
					}

					customers = append(customers[:delteIndex], customers[delteIndex+1:]...)
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
		case "5":
			{
				//保存客户功能
				fmt.Println("------------------保存客户信息开始------------------")
				//json序列化客户信息数据
				jsonData, err := json.Marshal(customers)
				ioutil.WriteFile("jsonCustomerdata.json", jsonData, 0666)
				if err != nil {
					fmt.Println("文件保存成功***")
				} else {
					fmt.Println("文件保存失败xxx")
				}

				fmt.Println("------------------保存客户信息结束------------------")

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
