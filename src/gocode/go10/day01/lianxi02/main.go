package main

import "fmt"

func main() {

	//var a1 int8 = 128 //int8类型是有符号的,-128到127,会溢出
	//fmt.Println(a1)

	////十进制
	//var a int = 10
	//fmt.Printf("%d\n", a) //10  %d表示十进制
	//fmt.Printf("%b\n", a) //1010  %b表示二进制
	//
	////八进制 以0开头
	//var a1 = 077
	//fmt.Printf("%o\n ", a1) //77
	////十六进制以0xx开头
	//var b = 0xff
	//fmt.Printf("%x\n", b) //ff
	//fmt.Printf("%X\n", b) //FF
	//
	////十进制转换
	//var c1 int = 10
	//fmt.Printf("%d\n", c1) //10   %d十进制
	//fmt.Printf("%b\n", c1) //1010   %b二进制
	//fmt.Printf("%o\n", c1) //12  %o八进制
	//fmt.Printf("%x\n", c1) //a  %x十六进制
	//
	////八进制转换
	//var c2 int = 020
	//fmt.Printf("%d\n", c2) //16  %d十进制
	//fmt.Printf("%b\n", c2) //10000   %b二进制
	//fmt.Printf("%o\n", c2) //20  %o八进制
	//fmt.Printf("%x\n", c2) //10 %x十六进制
	//
	////十六进制转换
	//var c3 = 0x12
	//fmt.Printf("%d\n", c3) //18  %d十进制
	//fmt.Printf("%b\n", c3) //10010   %b二进制
	//fmt.Printf("%o\n", c3) //22  %o八进制
	//fmt.Printf("%x\n", c3) //12 %x十六进制

	////打印数据类型
	//var a = 20
	//var b = "daerwen"
	//fmt.Println("a的类型是: ", a, reflect.TypeOf(a))
	//fmt.Println("b的类型是: ", b, reflect.TypeOf(b))

	////float 32单精度浮点型
	//var f1 float32
	//f1 = 3.1234567890123456789
	//fmt.Println("float 32 单精度浮点型: ", f1, reflect.TypeOf(f1))
	//
	////float 64 双精度浮点型
	//var f2 float64
	//f2 = 3.1234567890123456789
	//fmt.Println("float 64 双精度浮点型: ", f2, reflect.TypeOf(f2))
	//
	////根据平台默认是float64 双精度浮点型
	//f3 := 3.1234567890123456789
	//fmt.Println("默认的浮点型 :", f3, reflect.TypeOf(f3))

	////科学计数法
	//var f1 = 2e10
	//fmt.Println("即使是整数用科学计数法表示也是浮点型: ", f1, reflect.TypeOf(f1))
	//
	//var f2 = 2e-2
	//fmt.Println(f2, reflect.TypeOf(f2))

	////布尔类型 true false
	//var s bool //声明一个bool类型
	//s = 2 > 1  //2大于1为 true
	//fmt.Println(s, reflect.TypeOf(s))
	//
	//fmt.Println(1 > 2) //比较运算的结果是一个bool值
	//fmt.Println(5 == 5)
	//fmt.Println(5 != 6)
	////fmt.Println(5 == "5")  数字不能=字符串
	//
	////var name = "daerwen"
	////var name1 = name
	////fmt.Println(name1, reflect.TypeOf(name1))
	//
	//var name = "daerwen"
	//var name1 = name = "dafenqi"
	//fmt.Println(name1, reflect.TypeOf(name1))

	////字符串
	//var s = "wellcome to go lang!!!"
	//fmt.Println(s)
	//
	////字符串拼接
	//var s1 = "hello " + s
	//fmt.Println(s1)
	//
	////索引 字符串索引
	////获取g 索引从0开始
	//fmt.Println(s[12]) //默认是二进制的值
	////强转为字符串
	//fmt.Println(string(s[12]))
	//
	////切片 取连续的字符串 字符串[开始索引:结束索引]
	////切片取值slice[start:end] 取出的元素数量为: 开始位置-结束位置
	////取出well
	//fmt.Println(s[0:3]) //取出的值为:wel  原因顾头不顾尾,左闭右开
	//fmt.Println(s[0:4]) //取出完整的well
	//
	////切片的缺省状态
	////取出Wellcome以后的值
	//fmt.Println(s[9:])
	//
	////取出go之前的值
	//fmt.Println(s[:12])

	////转义字符
	////赋予某些符号以特殊功能
	////取消某些符号的特殊功能
	//fmt.Println("hello \nworld") //\n换行符
	//
	////转移特殊符号
	//fmt.Println("my name is \"daerwen\"")
	//
	////取消某些特殊符号的功能
	//fmt.Println("D:\\new\\data\\2022_07")

	////多行字符串
	//tanshi := `
	//		《木兰花令拟古决绝词》
	//				——纳兰容若
	//	人生若只如初见，何事秋风悲画扇。
	//`
	//fmt.Println(tanshi)

	////字符串常用方法
	//
	////求字符串长度
	//s := "daerwen,bijiAsuo,jialiluo"
	//fmt.Println(len(s))
	//
	////全部转大写
	//s1 := strings.ToUpper(s)
	//fmt.Println(s1)
	//
	////全部转小写
	//s2 := strings.ToLower(s)
	//fmt.Println(s2)
	//
	////去除字符串两端匹配的内容(比如空格)
	//s3 := "     Derwen     "
	//fmt.Println(strings.Trim(s3, " "))
	//
	////将字符串分割成数组
	//fmt.Println(strings.Split(s, ""))
	//fmt.Println(strings.Split(s, " "))
	//fmt.Println(strings.Split(s, ","))

	////将数组拼接成字符串
	//var citys1 = "北京-广州-上海-深圳"
	//var cityArr = strings.Split(citys1, "-")
	//fmt.Println(cityArr)
	//var citynew = strings.Join(cityArr, ";")
	//fmt.Println(citynew)

	////查询方法 strings.Index
	//var a1 = "daerwen reboot bijiasuo"
	//var strcd = len(a1)
	//fmt.Println(strcd)
	//var index = strings.Index(a1, "daerwen")
	//var index1 = strings.Index(a1, "bijiasuo")
	//fmt.Println(index)
	//fmt.Println(index1)
	//
	//var index2 = strings.LastIndex(a1, "reboot")
	//var index3 = strings.LastIndex(a1, "bijiasuo")
	////匹配到返回0,未匹配到返回-1
	//var index4 = strings.LastIndex(a1, "dafenqi")
	//fmt.Println(index2)
	//fmt.Println(index3)
	//fmt.Println(index4)

	////前缀判断,后缀判断,实现的依赖就是strings.Index 类型bool值
	//var a1 = "daerwen reboot bijiasuo"
	//fmt.Println(strings.HasPrefix(a1, "da"))
	//fmt.Println(strings.HasSuffix(a1, "suo"))
	////判断不是以xxx 结尾返回false
	//fmt.Println(strings.HasSuffix(a1, "sou"))
	//
	////判断是否包含 类型bool值
	//var a2 = "daerwen reboot bijiasuo"
	//fmt.Println(strings.Contains(a2, "wen"))
	////不包含返回false
	//fmt.Println(strings.Contains(a2, "beo"))

	//go语言不支持类型暗转换(隐式转换)
	//fmt.Println("100" > 20)

	//定义变量a 类型为int8
	//	var a int8
	//	a = 127
	//	fmt.Println(a, reflect.TypeOf(a))
	//	//类型转换type()  例如int64()  float32
	//	fmt.Println(int64(a), reflect.TypeOf(int64(a)))
	//	fmt.Println(float32(a), reflect.TypeOf(float32(a)))

	//var a int8
	//a = 100
	////字节转为字符,不能直接转为字符串
	//fmt.Println(string(a), reflect.TypeOf(string(a)))
	//
	////数字与字符串之间转换需要借助strconv库
	////把数字转为字符串
	//var age = 22
	//fmt.Println(age, reflect.TypeOf(age))
	////将int类型转为字符串
	//var agestr = strconv.Itoa(age)
	//fmt.Println(agestr, reflect.TypeOf(agestr))
	////将字符串数字类型转为字符串
	//var gao = "170"
	//fmt.Println(gao, reflect.TypeOf(gao))
	//var gao01, _ = strconv.Atoi(gao)
	//fmt.Println(gao01, reflect.TypeOf(gao01))
	//
	////浮点型字符串转为浮点型
	//x, _ := strconv.ParseFloat("3.14", 64)
	//fmt.Println(x, reflect.TypeOf(x))

	////运算符
	//fmt.Println("取余操作: ", 5%3)
	//
	////判断一个整形数字是奇数还是偶数
	//var x = 67
	//fmt.Println("偶数为ture 奇数为false")
	//fmt.Println("判断是奇数还是偶数: ", x%2 == 0)
	//fmt.Println("判断是奇数还是偶数: ", 20%2 == 0)
	//
	////自加操作
	////var a = 20
	////a = a + 1
	////fmt.Println(a)
	//var a = 20
	//a++
	//fmt.Println(a)
	//
	//var a1 = 30
	//a1 += 1
	//fmt.Println(a1)
	//
	////注意:不能写成++a 或 --a 必须放在右边使用
	//
	////自减操作
	////var b = 22
	////b = b - 1
	////fmt.Println(b)
	//
	//var b = 22
	//b--
	//fmt.Println(b)
	//
	//var b1 = 32
	//b1 -= 1
	//fmt.Println(b1)
	//
	////重新赋值内存地址不变,存储的值会重新赋值
	////&取地址操作符
	//var c = 10
	//fmt.Println(c, &c) //10 0x140000aa008
	//c = 100
	//fmt.Println(c, &c) //100 0x140000aa008

	////关系运算符  true false
	//fmt.Println(2 == 2)
	//fmt.Println(3 >= 1)
	//fmt.Println(5 != 6)
	//fmt.Println(7 <= 5)

	////逻辑运算符
	//fmt.Println(2 >= 1 && 1 == 1)
	//fmt.Println(2 <= 1 && 1 == 1)
	//
	//fmt.Println(2 == 2 || 2 <= 1)
	//fmt.Println(2 <= 1 || 1 != 1)
	//
	////用户名为jieke 或者大于18岁
	//var name = "jieke"
	//var age = 16
	//var ret = name == "jieke" || !(age > 18)
	//fmt.Println(ret)

	////运算符优先级
	//var a, b, c, d = 8, 6, 4, 2
	////先算乘除再算加减
	//ret := a + b*c/d
	//fmt.Println(ret)
	//
	//x := 10
	//y := 1
	////先算括号里的,然后再乘除在加减 ,最后的值 x+= 就是X加上值 所以是26
	//x += 5*(1+2) + y
	//fmt.Println(x) //26
	//
	//z := 1+2 > 3 || 1 == 1*5
	//fmt.Println(z)

	//name := "jieke"
	//age := 24
	//fmt.Print(name)
	//fmt.Print(age)
	//fmt.Println("hello go\n")
	//fmt.Println(name)
	//fmt.Println(age)

	//name := "jieke"
	//age := 24
	//isMarr := false
	//salary := 30000.549
	//fmt.Printf("姓名: %s  年龄: %d   婚否: %t  薪资: %.2f\n", name, age, isMarr, salary)
	//fmt.Printf("姓名: %v  年龄: %v   婚否: %v  薪资: %v\n", name, age, isMarr, salary)
	//fmt.Printf("姓名: %#v  年龄: %#v   婚否: %#v  薪资: %#v\n", name, age, isMarr, salary)

	////整形和浮点型
	//fmt.Printf("%b\n", 12)       //二进制表示:1100
	//fmt.Printf("%d\n", 12)       //十进制表示:12
	//fmt.Printf("%o\n", 12)       //八进制表示:14
	//fmt.Printf("%x\n", 12)       //十六进制表示:c
	//fmt.Printf("%X\n", 12)       //十六进制表示C
	//fmt.Printf("%f\n", 3.1415)   //有小数点而无指数: 3.141500
	//fmt.Printf("%.3f\n", 3.1415) //3.142
	//fmt.Printf("%e\n", 1000.25)  //科学计数法1.000250e+03,默认精度为6
	//
	////设置宽度
	//fmt.Printf("学号: %s 姓名: %-20s 平均成绩: %-4d\n", "1001", "jieke", 100)
	//fmt.Printf("学号: %s 姓名: %-20s 平均成绩: %-4d\n", "1002", "dijiaaotemanjiekeaoteman", 98)
	//fmt.Printf("学号: %s 姓名: %-20s 平均成绩: %-4d\n", "1003", "x", 78)

	//name := "jieke"
	//age := 22
	//isMarr := false
	//salary := 30000.549
	//info := fmt.Sprintf("姓名: %s 年龄: %d 婚否: %t 薪资: %.2f\n", name, age, isMarr, salary)
	//fmt.Println(info)

	//Scan
	//var (
	//	name   string
	//	age    int
	//	isMarr bool
	//)
	//fmt.Scan(&name, &age, &isMarr) //输入的类型不一致,按默认值
	//fmt.Printf("扫描结果 name:%s age:%d marride:%t\t", name, age, isMarr)

	//Scanf
	//var (10
	//	name   string
	//	age    int
	//	isMarr bool
	//)
	//fmt.Scanf("姓名是: %s 年龄是: %d 婚否: %t", &name, &age, &isMarr) //输入的类型不一致,按默认值
	//fmt.Printf("扫描结果 姓名:%s 年龄:%d 婚否:%t", name, age, isMarr)

	//var a, b int
	//fmt.Scanf("%d+%d", &a, &b)
	//fmt.Println(a + b)

	const (
		n1 = iota //0
		n2        //1
		_         //跳过 匿名占位
		n4        //3
	)
	fmt.Println(n1, n2, n4)

	const (
		a    = iota               //0
		b                         //1
		_                         //2  跳过匿名占位
		c, d = iota + 1, iota + 2 //c =3+1 =4  d= 3+ 2 =5
		e    = iota               //4
	)
	fmt.Println(a, b, c, d, e) //0 1 4 5 4

}
