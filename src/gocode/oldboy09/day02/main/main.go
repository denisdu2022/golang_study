package main

func main() {
	//十进制转换
	//var a int = 10
	//fmt.Printf("%d \n", a) //10   占位符%d表示十进制
	//fmt.Printf("%b \n", a) //1010   占位符%b表示二进制
	//fmt.Printf("%o \n", a) //12   占位符%o表示八进制
	//fmt.Printf("%x \n", a) //a   占位符%x表示十六进制

	//八进制转换
	//var b int = 020
	//fmt.Printf("%o \n", b) //20
	//fmt.Printf("%d \n", b) //16
	//fmt.Printf("%x \n", b) //10
	//fmt.Printf("%b \n", b) //10000

	//十六进制转换
	//var c = 0x12
	//fmt.Printf("%d \n", c) //18
	//fmt.Printf("%o \n", c) //22
	//fmt.Printf("%x \n", c) //12
	//fmt.Printf("%b \n", c) //10010

	//var f1 float32 = 3.14
	//fmt.Println(f1, reflect.TypeOf(f1))
	//
	//var f2 float64 = 3.14
	//fmt.Println(f2, reflect.TypeOf(f2))
	//
	//var f3 = 3.14
	//fmt.Println(f3, reflect.TypeOf(f3)) //根据当前平台是64

	//var f1 = 2e10 //即使是整数用科学计数表示也是浮点型
	//fmt.Println(f1, reflect.TypeOf(f1))
	//
	//var f2 = 2e-2
	//fmt.Println(f2, reflect.TypeOf(f2))

	//fmt.Println(1 == 1)
	//fmt.Println(2 < 1)
	//fmt.Println(5 > 3)
	//b := 3 > 1
	//fmt.Println(b, reflect.TypeOf(b))

	//fmt.Println('a')                //单引号''是字符
	//fmt.Println("heelo world")      //字符串
	//fmt.Println("hello" + "world!") //字符串拼接

	//var s = "hello world reboot"
	//fmt.Println(s)
	//
	////1.索引取值 slice[index]
	//a := s[2]
	//fmt.Println(string(a))
	//
	////2.切片取值slice[start:end] ,取出的元素数量为:结束位置 -开始位置;
	//b1 := s[3:6] //顾头不顾尾
	//fmt.Println(b1)
	//b2 := s[0:] //当缺省结束位置时,表示从开始位置到整个连续区域末尾
	//fmt.Println(b2)
	//b3 := s[:8] //当缺省开始位置时,表示从连续区域开头到结束位置
	//fmt.Println(b3)
	//
	////3.字符串拼接
	//var s1 = "hello"
	//var s2 = "world reboot"
	//var s3 = s1 + s2 //拼接成一个新的字符串
	//fmt.Println(s3)

	/*	var s1 = "hi reboot\nhi world" //换行符
		fmt.Println(s1)
		var s2 = "my is \"daerwen\"" //转义特殊符号
		fmt.Println(s2)
		var s3 = "D:\\new\\golang\\go.exe"
		fmt.Println(s3)*/

	//多行字符串
	//s1 := `春日
	//胜日寻芳泗水滨
	//无边光景一时新
	//等闲识得东风面
	//万紫千红总是春
	//`
	//fmt.Println(s1)

	////求字符串长度
	//s := "daerwen,bijiasuo,jialiluo"
	//fmt.Println(len(s))
	//
	////全部转大写
	//s1 := strings.ToUpper("Derwen")
	////全部转小写
	//s2 := strings.ToLower("BjiSuo")
	//
	//fmt.Println("s1 转换后是:", s1)
	//fmt.Println("s2 转换后是:", s2)
	//
	////去除字符串两端匹配的内容(比如空格)
	//s3 := strings.Trim("    Daerwen   ", " ")
	//fmt.Println(s3, len(s3))
	//
	////将字符串分割成数组
	//var ret = strings.Split(s, ",") //以,分割
	//fmt.Println(ret, reflect.TypeOf(ret))
	//
	////将数组拼接成字符串
	//var pjret = strings.Join(ret, "~") //以~拼接
	//fmt.Println(pjret, reflect.TypeOf(pjret))
	//
	////以...开头,以...结尾
	//var index = strings.Index(s, "daerwen")
	//fmt.Println(index, reflect.TypeOf(index))
	////匹配到返回0,未匹配到返回-1
	//var index2 = strings.LastIndex(s, "reboot")
	//fmt.Println(index2, reflect.TypeOf(index2))
	//
	////前缀判断,后缀判断,实现的依赖就是strings.Index 类型bool值
	//fmt.Println(strings.HasPrefix(s, "da"))
	//fmt.Println(strings.HasSuffix(s, "la"))
	//
	////包含  类型bool值
	//fmt.Println(strings.Contains(s, "a"))
	//fmt.Println(strings.Contains(s, "1"))

	//var a int8
	//a = 100
	//fmt.Println(reflect.TypeOf(int64(a)))
	//fmt.Println(reflect.TypeOf(float64(a)))
	////字节转换为字符,不能直接转为字符串
	//fmt.Println(string(a), reflect.TypeOf(string(a)))
	//
	////数字与字符串之间的转换需要借助strconv库
	//
	//x := strconv.Itoa(112)
	//fmt.Println(x, reflect.TypeOf(x))
	//y, _ := strconv.Atoi("1200")
	//fmt.Println(y, reflect.TypeOf(y))
	//
	////浮点型字符串转为浮点型
	//
	//z, _ := strconv.ParseFloat("3.14", 64)
	//fmt.Println(z, reflect.TypeOf(z))

	//a := 1
	//a++ //注意:不能写成++a 或 --a 必须放在右边使用
	//
	//fmt.Println(a) //a = 2

}
