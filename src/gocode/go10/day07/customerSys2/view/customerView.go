package view

import (
	"customerSys2/service"
	"fmt"
	"os"
)

func CustomerView() {

	//实例化客户信息
	cs := service.CustomerSys{}

	//初始化客户信息数据
	cs.InitCusomers()
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
