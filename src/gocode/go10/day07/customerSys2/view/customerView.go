package view

import (
	"fmt"
	"os"
)

func CustomerView() {
	for true {
		var choice int
		fmt.Println(`
------------------客户信息管理系统------------------
			1. 添加客户信息
			2. 查看客户信息
			3. 更新客户信息
			4. 删除客户信息
			5. 保存客户信息
			6. 退出客户信息管理系统
------------------------------------------------------
`)
		fmt.Print("请输入您的选择[1-6]: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			{
				//添加客户功能
			}
		case 2:
			{
				//查看客户功能
			}
		case 3:
			{
				//更新客户功能
			}
		case 4:
			{
				//删除客户功能
			}
		case 5:
			{
				//保存客户功能
			}
		case 6:
			{
				//退出客户信息管理系统
				fmt.Println("您已退出客户信息管理系统,欢迎再次使用...")
				os.Exit(0)
			}
		default:
			fmt.Println("您输入的选择数值不合法,请您重新输入....")
			continue
		}
	}

}
