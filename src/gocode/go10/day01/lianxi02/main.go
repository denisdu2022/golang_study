package main

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

}
