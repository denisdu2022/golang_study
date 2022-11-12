package main

import (
	"fmt"
	"os"
	"strconv"
)

var customers []map[string]string
var customerID int

func main() {

	for true {
		fmt.Printf("\033[1;30;42m%s\033[0m\n", `
----------------客户信息管理系统--------------
		   1、添加客户
		   2、查看客户
		   3、更新客户
		   4、删除客户
           5、保存
		   6、退出
-------------------------------------------
`)

		var choice int
		fmt.Print("请输入您的选择：")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// 添加客户
			// 引导用户输入

			// 引导用户输入学号和姓名
			fmt.Printf("\033[1;35;40m%s\033[0m\n", "---------------------------添加客户开始-----------------------------")
			// 引导用户输入
			var name string
			fmt.Print("请输入客户姓名：")
			fmt.Scan(&name)

			var gender string
			fmt.Print("请输入客户性别：")
			fmt.Scan(&gender)

			var age string
			fmt.Print("请输入客户年龄：")
			fmt.Scan(&age)

			var email string
			fmt.Print("请输入客户邮箱：")
			fmt.Scan(&email)
			customerID++
			var newCustomer = map[string]string{
				"cid":    strconv.Itoa(customerID),
				"name":   name,
				"age":    age,
				"gender": gender,
				"email":  email,
			}
			customers = append(customers, newCustomer)
			fmt.Printf("\033[1;35;40m%s\033[0m\n", "---------------------------添加客户成功-----------------------------")

		case 2:
			fmt.Printf("\033[1;32;40m%s\033[0m\n", "----------------------------------客户列表开始----------------------------")

			for _, customerMap := range customers {
				fmt.Printf("客户编号：%s 姓名：%-6s 性别：%-6s 年龄：%-6s 邮箱：%-6s\n", customerMap["cid"], customerMap["name"], customerMap["gender"], customerMap["age"], customerMap["email"])
			}
			fmt.Printf("\033[1;32;40m%s\033[0m\n", "----------------------------------客户列表结束----------------------------")

		case 3:

			for true {
				//引导用户输入一个客户编号
				var cid string
				fmt.Print("请输入修改的客户编号：")
				fmt.Scan(&cid)
				//  判断 cid的客户索引值，-1代表不存在
				var updateIndex = -1
				for index, customer := range customers {
					if customer["cid"] == cid {
						updateIndex = index
					}
				}

				if updateIndex == -1 {
					fmt.Println("该编号不存在")
					continue
				}
				// 获取更细的客户map对象
				updateCustomer := customers[updateIndex]
				// 引导用户输入
				var name string
				fmt.Printf("请输入客户姓名(%s)：", updateCustomer["name"])
				fmt.Scanln(&name)

				var gender string
				fmt.Printf("请输入客户性别(%s)：", updateCustomer["gender"])
				fmt.Scanln(&gender)

				var age string
				fmt.Printf("请输入客户年龄(%s)：", updateCustomer["age"])
				fmt.Scanln(&age)

				var email string
				fmt.Printf("请输入客户邮箱(%s)：", updateCustomer["email"])
				fmt.Scanln(&email)

				if name != "" {
					updateCustomer["name"] = name
				}
				if gender != "" {
					updateCustomer["gender"] = gender
				}
				if age != "" {
					updateCustomer["age"] = age
				}
				if email != "" {
					updateCustomer["email"] = email
				}
				break

			}

		case 4:
			for true {
				//引导用户输入一个客户编号
				var cid string
				fmt.Print("请输入删除的客户编号：")
				fmt.Scan(&cid)
				//  判断 cid的客户索引值，-1代表不存在
				var deleteIndex = -1
				for index, customer := range customers {
					if customer["cid"] == cid {
						deleteIndex = index
					}
				}
				if deleteIndex == -1 {
					fmt.Println("该删除编号不存在")
					continue
				}

				// 删除  按索引删除

				customers = append(customers[:deleteIndex], customers[deleteIndex+1:]...)
				fmt.Println(customers)
				fmt.Printf("\033[1;31;40m%s\033[0m\n", "---------------------------删除客户完成----------------------")
				break

			}

		case 6:
			os.Exit(0)
		}

		fmt.Println("over!")

	}

}
