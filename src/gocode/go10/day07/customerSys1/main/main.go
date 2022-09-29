package main

import (
	"encoding/json"
	"fmt"
	"gocode/go10/day07/customerSys1/model"
	"gocode/go10/day07/customerSys1/service"
	"io/ioutil"
	"os"
)

func main() {

	//初始化保存数据
	//读取保存数据的文件
	customerJsonBytes, err := ioutil.ReadFile("db/customerData.json")
	if err != nil {
		fmt.Println("初始化客户信息失败,错误原因: ", err)
	} else {
		fmt.Println("初始化客户信息成功........")
	}
	//定义反序列化接收变量
	var customers []model.Customer
	//json反序列化
	json.Unmarshal(customerJsonBytes, &customers)

	//定义初始化自增ID initCustomerID为0
	var initCustomerID = 0
	//当customers的长度不等于0的时候让初始的自增IDinitCustomerID == customers的索引是customers的长度-1.的键Cid
	if len(customers) != 0 {
		initCustomerID = customers[len(customers)-1].Cid
	}

	//实例化CustomerService
	cs := service.CustomerService{
		Customers:  customers,
		CustomerID: initCustomerID,
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
				cs.AddCustomer()
			}
		case "2":
			{
				cs.ListCustomer()
			}
		case "3":
			{
				cs.UpdataCustomer()
			}
		case "4":
			{
				cs.DelteCustomer()
			}
		case "5":
			{
				cs.SaveCustomer()
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
