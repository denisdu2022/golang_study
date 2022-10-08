package view

import (
	"customerSys2/model"
	"customerSys2/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func CustomerView() {
	//初始化保存数据
	//读取保存数据的文件
	customerJsonBytes, err := ioutil.ReadFile("db/customerJsonData.json")
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

	//实例化客户信息
	cs := service.CustomerSys{
		Customer:   customers,
		CustomerId: initCustomerID,
	}

	////初始化客户信息数据
	//cs.InitCusomers()
	for true {
		var choice int
		fmt.Println(`
-------------------客户信息管理系统-------------------
			1. 添加客户信息
			2. 查看客户信息
			3. 更新客户信息
			4. 删除客户信息
			5. 保存客户信息
			6. 退出客户信息管理系统
---------------------------------------------------------
`)
		//引导用户输入选择客户信息管理系统功能
		fmt.Print("请输入您的选择[1-6]: ")
		fmt.Scanln(&choice)
		//客户信息管理系统功能选择
		switch choice {
		case 1:
			{
				//添加客户信息
				cs.AddCustome()
			}
		case 2:
			{
				//查看客户信息
				cs.ListCustome()
			}
		case 3:
			{
				//更新客户信息
				cs.UpdataCustome()
			}
		case 4:
			{
				//删除客户信息
				cs.DelCustome()
			}
		case 5:
			{
				//保存客户信息
				cs.SaveCustome()
			}
		case 6:
			{
				//退出客户信息管理系统
				fmt.Println("您已退出客户信息管理系统,欢迎再次使用!!!")
				os.Exit(0)
			}
		default:
			//非法输入
			fmt.Println("您输入的数值不合法,请您重新输入...")
			continue
		}
	}
}
