package main

import "fmt"

func main() {

	////交换x和y两个变量的值
	//x := 10
	//y := 20
	//x, y = y, x
	//fmt.Println(x, y)

	////将字符串转为整形
	//agestr := "18"
	//ageint, _ := strconv.Atoi(agestr)
	//fmt.Println(ageint, reflect.TypeOf(ageint))
	//
	////将整形转为字符串
	//sint := 22
	//sstr := strconv.Itoa(sint)
	//fmt.Println(sstr, reflect.TypeOf(sstr))
	//
	////将浮点型字符串转为浮点型
	//spi := "3.1415926"
	//spif, _ := strconv.ParseFloat(spi, 64)
	//fmt.Println(spif, reflect.TypeOf(spif))1992

	////输入生日,显示格式为年-月-日
	//var birth string
	//fmt.Println("请输入你的生日: 格式为年 月 日")
	//fmt.Scan(&birth)
	//newBirth := strings.Split(birth, " ")
	//fmt.Printf("你的生日是: %s年 %s月 %s日", newBirth[0], newBirth[1], newBirth[2])

	////引导用户输入姓名,判断是否是以小写的a,或者大写的A开头,是返回TRUE,否则是FALSE
	//var name string
	//println("请输入你的名字: ")
	//fmt.Scan(&name)
	//fmt.Println(name)
	//pd := strings.HasPrefix(name, "a") || strings.HasPrefix(name, "A")
	//fmt.Printf("是否是以小写的a,或者大写的A开头,是返回TRUE,否则是FALSE: %t\n", pd)

	////输入用户名和密码为,if判断 name = "daerwen" ,pwd = 123
	//var name string
	//var pwd int
	//fmt.Println("请输入你的用户名和密码:")
	//fmt.Scan(&name, &pwd)
	//if name == "daerwen" && pwd == 123 {
	//	fmt.Println("登录成功")
	//} else {
	//	fmt.Println("登录失败")
	//}

	////输入生日判断星座
	//var birth string
	//fmt.Println("请输入你的生日,格式: 1997-04-03 ")
	//fmt.Scan(&birth)
	//fmt.Println(birth, reflect.TypeOf(birth))
	//birthf := strings.Split(birth, "-")
	//fmt.Println(birthf)
	//month := birthf[1]
	//dayth := birthf[2]
	//mdhb := month + "." + dayth
	//birthff, _ := strconv.ParseFloat(mdhb, 64)
	//fmt.Println(birthff)
	//
	//if birthff >= 3.21 && birthff <= 4.19 {
	//	fmt.Println("你是白羊座")
	//} else if birthff >= 4.20 && birthff <= 5.20 {
	//	fmt.Println("你是金牛座")
	//} else if birthff >= 5.21 && birthff <= 6.21 {
	//	fmt.Println("你是双子座")
	//} else if birthff >= 6.22 && birthff <= 7.22 {
	//	fmt.Println("你是巨蟹座")
	//} else if birthff >= 7.23 && birthff <= 8.22 {
	//	fmt.Println("你是狮子座")
	//} else if birthff >= 9.23 && birthff <= 10.22 {
	//	fmt.Println("你是处女座")
	//} else if birthff >= 11.23 && birthff <= 12.21 {
	//	fmt.Println("你是天秤座")
	//} else if birthff >= 12.22 && birthff <= 1.19 {
	//	fmt.Println("你是天蝎座")
	//} else if birthff >= 1.20 && birthff <= 2.18 {
	//	fmt.Println("你是射手座")
	//} else if birthff >= 2.19 && birthff <= 3.20 {
	//	fmt.Println("你是水瓶座")
	//}
	var sum = 0

	for i := 1; i < 100; i++ {

		if i == 88 {
			continue
		}

		if i%2 != 0 {
			sum += i
		} else {
			sum -= i
		}

		fmt.Println(sum)

	}

}
