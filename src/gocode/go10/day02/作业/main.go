package main

func main() {

	////1.一个整数判断是奇数还是偶数
	////奇数是false,偶数数true
	//var x = 67
	//fmt.Println("67 这是是偶数还是奇数: ", x%2 == 0)
	//fmt.Println("66 这是是偶数还是奇数: ", 66%2 == 0)

	//2
	//var a int8 = 100
	//var b int16 = 200
	////fmt.Println(a + b)
	//fmt.Println(int16(a) + b)

	////3
	//cmd := ""
	//cmd += "l"
	//cmd += "s"
	//fmt.Println(cmd)

	//4
	//var birth string
	//fmt.Println("请输入您的生日(格式为:199-3-16): ")
	//fmt.Scan(&birth)
	//nBirth := strings.Split(birth, "-")
	//fmt.Println(nBirth)
	//year := nBirth[0]
	//month := nBirth[1]
	//day := nBirth[2]
	//
	//fmt.Print("您的生日是:", year+"年-", month+"月-", day+"日")

	//5
	//var user string
	//fmt.Println("请输入您的名字: ")
	//fmt.Scan(&user)
	//
	//fmt.Println(strings.HasPrefix(user, "a") || strings.HasPrefix(user, "A"))

	//var user string
	//fmt.Println("请输入您的名字: ")
	//fmt.Scan(&user)
	//nUser := strings.ToUpper(user)
	//
	//fmt.Println(strings.HasPrefix(nUser, "A"))

	//6
	//var name,age = "yuan" ,22
	//fmt.Println(name == "yuan" && age = 23)

	//7
	//var x, y = 10, 5
	//x += y
	//fmt.Println(x)

	//8
	//y := strconv.Itoa(100)
	//
	//fmt.Println(y, reflect.TypeOf(y))

	//9
	//f := 3.1415926
	//fmt.Printf("这是%T\n", f)   //打印对应值的类型
	//fmt.Printf("这是%#v\n", f)  //相应值的go语法表示
	//fmt.Printf("这是%.2f\n", f) //小数点指数,.2f是小数点后两位

	//11
	//f1 := "3.24"
	//f2, _ := strconv.ParseFloat(f1, 64)
	//fmt.Println(f2, reflect.TypeOf(f2))

	//12
	//var a int8 = 100      //定义变量a ,为int 8类型,并赋值100
	//b := a                //定义变量b 赋值b的值为a ,所以b是100
	//a++                   //a自加1 是101
	//fmt.Println("a :", a) //101
	//fmt.Println("b :", b) //100

	//13
	//var year string
	//fmt.Println("请您输入一个年份: ")
	//fmt.Scan(&year)
	//
	//nYear, _ := strconv.Atoi(year)
	//
	//if nYear%4 == 0 || nYear%400 == 0 {
	//	fmt.Println("您输入的是闰年")
	//} else {
	//	fmt.Println("您输入的是平年")
	//}

	//16
	//var sum = 0
	//
	//for i := 1; i <= 10; i++ {
	//	var c = 1
	//	for b := 1; b <= i; b++ {
	//		c *= b
	//	}
	//	sum += c
	//}
	//fmt.Println(sum)

}
