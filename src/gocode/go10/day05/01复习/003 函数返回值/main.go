package main

import "fmt"

////有返回值的函数
//func add(a, b int) int {
//	c := a + b
//	return c //函数的终止语句
//
//}

////多个返回值
//func login(user, pwd string) (bool, string) {
//	if user == "root" && pwd == "123" {
//		//登录成功
//		return true, user
//	} else {
//		//登录失败
//		return false, user
//	}
//}

//返回值命名
func add(s ...int) (z int) {

	for _, v := range s {
		z += v
	}
	return

}

func main() {

	////调用函数的返回值
	////如果是无返回值的函数调用是不可以赋值给一个变量的
	//ret := add(1, 2)
	//fmt.Println(ret)
	//ret = add(ret, 100) //把函数返回值在作为参数传给add的参数
	//fmt.Println(ret)    //103

	////调用多个返回值的函数
	////login1, ok := login("root", "123")
	////fmt.Println(login1, ok)
	//ok, user := login("root", "123")
	//if ok {
	//	fmt.Println("登录成功,用户名:", user)
	//} else {
	//	fmt.Println("登录失败")
	//}

	//返回值重命名
	sum := add(1, 2, 3, 4, 5) //15
	fmt.Println(sum)
}
