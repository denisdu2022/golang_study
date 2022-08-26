package main

////匿名函数类型 func()
//func foo() {
//	var bar = func() {
//		fmt.Println("bar函数的功能")
//	}
//	bar()
//	fmt.Println(reflect.TypeOf(bar))  //func()
//}
//
////匿名函数类型 func()带参数
//func foo() {
//	var bar func(int, int)
//	bar = func(x, y int) {
//		fmt.Println(x + y)
//		fmt.Println("bar函数的功能")
//	}
//	bar(12, 13)
//	fmt.Println(reflect.TypeOf(bar)) //func(int int)
//}

func main() {
	////匿名函数类型
	//foo()
	//
	////声明匿名函数并赋值调用
	//var f = func() {
	//	fmt.Println("匿名函数...")
	//}
	//f()
	//
	////直接声明并调用匿名函数
	//(func() {
	//	fmt.Println("直接声明并调用匿名函数")
	//})()
	//
	////匿名函数带参数
	//(func(x, y int) { //函数里的形参
	//	fmt.Println("x+y的和: ", x+y)
	//})(12, 13) //调用执行的实参

}
