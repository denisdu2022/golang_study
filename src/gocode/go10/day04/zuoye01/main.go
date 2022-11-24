package main

import (
	"fmt"
)

func main() {

	var xuanZe = 0
	var names, gender, ages, email string
	cmsMap := []map[string]string{}

	for true {
		wel := `
----------------客户信息管理系统----------------
		1, 添加客户
		2, 查看客户
		3, 更新客户
		4, 删除客户
		5, 退出
----------------------------------------------
`
		fmt.Println(wel)
		fmt.Println("请输入您的选择: ")
		fmt.Scan(&xuanZe)
		switch xuanZe {
		case 1:
			fmt.Println("---------------添加客户开始---------------")
			fmt.Println("请输入客户姓名:")
			fmt.Scan(&names)
			fmt.Println("请输入客户性别:")
			fmt.Scan(&gender)
			fmt.Println("请输入客户年龄:")
			fmt.Scan(&ages)
			fmt.Println("请输入客户邮箱:")
			fmt.Scan(&email)
			fmt.Println("---------------添加客户完成---------------")
			cmsMap = append(cmsMap, map[string]string{"name": names, "gender": gender, "age": ages, "email": email})
			//fmt.Println(cmsMap)
		case 2:
			fmt.Println("---------------客户列表开始---------------")
			for _, newCmsMap := range cmsMap {

				fmt.Printf("姓名是:%-8s 性别是:%-8s 年龄是:%-8s 邮箱是:%-8s", newCmsMap["name"], newCmsMap["gender"], newCmsMap["age"], newCmsMap["email"])
				fmt.Println()
				//fmt.Println(newCmsMap)
				//for _, v := range newCmsMap {
				//	//fmt.Println(cmsMapData, reflect.TypeOf(cmsMapData))
				//	//fmt.Println(v)
				//	//fmt.Printf("姓名是:%-8s,性别是:%8s,年龄是:%-8s,邮箱是:%-8s", )
				//}
			}
			fmt.Println("---------------客户列表结束---------------")
		case 3:
			fmt.Println("---------------客户修改开始---------------")
			fmt.Println("请输入修改的客户信息:")
			fmt.Println("请输入客户姓名:")
			fmt.Scan(&names)
			fmt.Println("请输入客户性别:")
			fmt.Scan(&gender)
			fmt.Println("请输入客户年龄:")
			fmt.Scan(&ages)
			fmt.Println("请输入客户邮箱:")
			fmt.Scan(&email)
			fmt.Println("---------------客户修改结束---------------")
			cmsMap = append(cmsMap, map[string]string{"name": names, "gender": gender, "age": ages, "email": email})
		case 4:
			fmt.Println("")
		case 5:
			fmt.Println("您已退出客户信息管理系统...")
			break
		}
	}

}
