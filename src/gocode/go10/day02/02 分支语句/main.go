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

}
