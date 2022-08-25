package main

////相加
//func add(x, y int) int {
//	return x + y
//}
//
////相乘
//func mul(x, y int) int {
//	return x * y
//}
//
////双值计算器
//func cal(f func(x, y int) int, x, y int) int {
//	return f(x, y)
//}

////相加函数
//func addCal(x int, y int) int {
//	return x + y
//}

////计算运行时间的函数,
//func timer(f func()) {
//	//定义初始时间戳
//	timeBefore := time.Now().Unix()
//	//定义了一个函数参数
//	f()
//	//定义结束时间戳
//	timeAfter := time.Now().Unix()
//	//开始减去结束 == 运行时间
//	fmt.Println("运行时间: ", timeBefore-timeAfter)
//}
//
////foo函数功能
//func foo() {
//	fmt.Println("foo 程序功能开始")
//	//使程序睡眠2s
//	time.Sleep(time.Second * 2)
//	fmt.Println("foo 程序功能结束")
//}
//
////bar函数功能
//func boo() {
//	fmt.Println("bar 程序功能开始")
//	time.Sleep(time.Second * 3)
//	fmt.Println("bar 程序功能结束")
//}

func main() {
	////调用双值计算函数
	//ret := cal(add, 13, 14)
	//fmt.Println(ret)
	//ret1 := cal(mul, 2, 4)
	//fmt.Println(ret1)

	//var a = addCal
	//a(2, 3)
	//fmt.Println(a)                      //0x107b4e0
	//fmt.Println(addCal)                 //0x107b4e0
	//fmt.Println(reflect.TypeOf(addCal)) //func(int, int) int

	////调用函数
	//timer(foo)
	//timer(boo)

}
