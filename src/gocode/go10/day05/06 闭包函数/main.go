package main

////闭包函数
//func foo() func() { //外层函数返回嵌套函数
//	var x = 100
//	var bar = func() { //嵌套函数
//		fmt.Println("bar功能开始: ", x) //嵌套函数必须引用外部非全局的局部自由变量
//	}
//	return bar
//
//}
//
////闭包函数,带参数
//func addSum() func(int) int {
//	sum := 15
//	return func(i int) int {
//		sum += i
//		return sum
//	}
//
//}

func main() {

	////调用闭包函数
	//var f = foo() //f是闭包函数
	//f()           //bar功能开始:  100  闭包函数使用完不会销毁
	//
	////调用带参数的闭包函数
	//f1 := addSum()(20)
	//fmt.Println(f1) //35
}
