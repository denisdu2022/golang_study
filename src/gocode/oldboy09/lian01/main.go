package main

func main() {
	//整形 int
	//var x = 100
	//fmt.Println(x, reflect.TypeOf(x)) //typeof 打印数据类型 //默认int类型

	//overflows
	//var y int8
	//y = 160
	//fmt.Println(y)

	////布尔类型: true false
	//var b bool
	//b = 2 > 1 //赋值存储以供后续使用
	//fmt.Println(b)

	////索引  字符串[索引]
	//s := "awellcome to golang!"
	//fmt.Println(s[4])         //默认是二进制的值
	//fmt.Println(string(s[4])) //强转为字符串

	////切片 取连续的字符串 字符串[开始索引:结束索引]
	//fmt.Println(s[1:9]) //取出Wellcome

	//转义符
	//功能1: 赋予某些普通符号以特殊功能
	//功能2: 取消某些特殊符号的特殊功能

	//go语言不支持类型暗转换(隐式转换)
	//fmt.Println("100" > 20)

	//var a int8
	//a = 127
	//fmt.Println(a, reflect.TypeOf(a))
	//
	//b := int64(a)
	//fmt.Println(b, reflect.TypeOf(b))

	//var age = 22
	//fmt.Println(age, reflect.TypeOf(age))
	////将int类型转为字符串的方法
	//var ageStr = strconv.Itoa(age)
	//fmt.Println(ageStr, reflect.TypeOf(ageStr))
	//
	////把字符串数字转为int类型
	//var money = "100"
	//moneyNum, err := strconv.Atoi(money)
	//fmt.Println("err: ", err)
	//fmt.Println("moneyNum: ", moneyNum)

	////把浮点型字符串转为浮点型
	//var f = "3.1415926"
	//f2, _ := strconv.ParseFloat(f, 64)
	//fmt.Println(f2, reflect.TypeOf(f2))

	//取余%
	//fmt.Println(5 % 3)
	//
	//var x = 67
	//fmt.Println(x%2 == 0)

	////重新赋值内存地址不变,存储的值会重新赋值
	//var a = 10
	//fmt.Println(a)
	//fmt.Println(&a)
	//a = 100
	//fmt.Println(&a)
	//fmt.Println(a)

	////关系运算符
	//fmt.Println(2 == 2)
	//fmt.Println(2 <= 3)

	//逻辑运算符

}
