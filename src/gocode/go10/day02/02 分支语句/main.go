package main

func main() {

	////01 顺序语句,从上到下依次执行
	//fmt.Println("世界你好 !!!!!")
	//fmt.Println("世界你好 !!!!!")
	//fmt.Println("世界你好 !!!!!")
	//fmt.Println("世界你好 !!!!!")
	//fmt.Println("世界你好 !!!!!")

	////02 单分支语句
	////语法
	//if bool表达式 { //花括号必须与表达式同行
	//	//在bool表达式为 true 时执行
	//}

	////单分支,判断输入的用户名密码是否正确
	//var username, userpwd string
	//
	//fmt.Println("请输入用户名: ")
	//fmt.Scanln(&username)
	//fmt.Println(username)
	//
	//fmt.Println("请输入密码: ")
	//fmt.Scanln(&userpwd)
	//fmt.Println(userpwd)
	//
	////判断用户名和密码是否为daerwen 和123
	//if username == "daerwen" && userpwd == "123" {
	//	fmt.Println("用户名密码输入正确,登录成功!!!")
	//}

	////03 双分支语句
	////满足年龄观看影片
	//var age int
	//
	//fmt.Scanln(&age)
	//
	//fmt.Println("请输入你的年龄: ")
	//if age >= 18 {
	//	fmt.Println("满足年龄限制,你可以观看. ")
	//} else {
	//	fmt.Println("很遗憾,年龄受限,不允许观看")
	//}

	////使用双分支完善上边的用户名密码输入
	//var username, userpwd string
	//
	//fmt.Println("请输入用户名: ")
	//fmt.Scanln(&username)
	//fmt.Println(username)
	//
	//fmt.Println("请输入密码: ")
	//fmt.Scanln(&userpwd)
	//fmt.Println(userpwd)
	//
	////判断用户名和密码是否为daerwen 和123
	//if username == "daerwen" && userpwd == "123" {
	//	fmt.Println("用户名密码输入正确,登录成功!!!")
	//} else {
	//	fmt.Println("用户名密码输入错误,登录失败!!!")
	//}

	////04 多分支语句
	////输入成绩,并判断 优秀,良好,及格,以及不及格
	//
	//var score int
	//fmt.Scanln(&score)
	//
	//if score > 0 && score <= 100 {
	//	if score > 90 {
	//		fmt.Println("成绩优秀 *****")
	//	} else if score > 70 {
	//		fmt.Println("成绩良好 ****")
	//	} else if score >= 60 {
	//		fmt.Println("成绩及格 ***")
	//	} else {
	//		fmt.Println("成绩不及格 ......")
	//	}
	//
	//} else {
	//	fmt.Println("输入非法数值!!!")
	//}

	//	s1 := `
	//			给出如下选项，并根据选项进行效果展示：
	//		    输入1：则输出"普通攻击"
	//		    输入2：则输出"超级攻击"
	//		    输入3：则输出"使用道具"
	//		    输入3：则输出"逃跑"
	//`
	//	fmt.Println(s1)
	//
	//	var gongji int
	//	fmt.Println("请输入你的选择")
	//	fmt.Scan(&gongji)
	//	fmt.Println(gongji, reflect.TypeOf(gongji))
	//
	//	switch gongji {
	//	case 1:
	//		fmt.Println("普通攻击")
	//	case 2:
	//		fmt.Println("超级攻击")
	//	case 3:
	//		fmt.Println("使用道具")
	//	case 4:
	//		fmt.Println("逃跑")
	//	default:
	//		fmt.Println("非法输入")
	//	}

	////switch 支持多条件匹配
	//var i1 int
	//fmt.Println("请输入你的选择: ")
	//fmt.Scan(&i1)
	//
	//switch i1 {
	//case 1, 2:
	//	fmt.Println("<<阿凡达>>")
	//case 3, 4:
	//	fmt.Println("阿凡提")
	//}

	////输入一个0-9之间的数判断奇数还是偶数
	//var jiO int
	//fmt.Println("请输入0-9之间的任意一个数字: ")
	//fmt.Scan(&jiO)
	//
	////switch多条件判断奇数偶数
	//switch jiO {
	//case 1, 3, 5, 7, 9:
	//	fmt.Println("输入的是奇数")
	//case 2, 4, 6, 8:
	//	fmt.Println("输入的是偶数")
	//default:
	//	fmt.Println("输入的是0")
	//}

	////分支嵌套
	//var userName, userPwd string
	//fmt.Println("请输入你的用户名: ")
	//fmt.Scan(&userName)
	//fmt.Println("请输入你的密码: ")
	//fmt.Scan(&userPwd)
	//
	////判断输入的用户名密码是否正确
	//if userName == "daerwen" && userPwd == "123" {
	//	fmt.Println("用户名密码正确,登录成功!!!")
	//	fmt.Println("**********************************************")
	//
	//	//输入成绩并判断
	//	var score int
	//	fmt.Println("请输入你的成绩: ")
	//	fmt.Scan(&score)
	//
	//	if score > 0 && score <= 100 {
	//		if score >= 90 {
	//			fmt.Println("成绩优秀!!!")
	//		} else if score >= 70 {
	//			fmt.Println("成绩良好!!!")
	//		} else if score >= 60 {
	//			fmt.Println("成绩合格!!!")
	//		} else if score < 60 {
	//			fmt.Println("成绩不合格")
	//		}
	//	} else {
	//		fmt.Println("输入的成绩分数是非法数值.....")
	//	}
	//} else {
	//	fmt.Println("用户名密码输入错误~~~")
	//}

}
