package main

import (
	"fmt"
	"reflect"
)

//定义一个空接口
type AnyKong interface {
}

//函数参数空接口
func foo(x interface{}) {
	fmt.Println(x, reflect.TypeOf(x))
}

func main() {

	var ak AnyKong  //空接口任意类型any
	ak = 10         //int类型
	fmt.Println(ak) //10
	ak = "hello!!!" //字符串类型
	fmt.Println(ak) //hello!!!

	//函数参数空接口
	foo(100)             //100 int
	foo("哈哈")            //哈哈 string
	foo([3]int{1, 2, 3}) //[1 2 3] [3]int  int数组
	foo([]int{4, 5, 6})  //[4 5 6] []int   int切片

	var m = map[string]interface{}{"name": "daerwen", "age": 23, "city": []string{"山东省", "北京市"}}
	fmt.Println(m) //map[age:23 city:[山东省 北京市] name:daerwen]

	//类型断言
	var x1 interface{}
	x1 = "hello"
	v, ok := x1.(string)
	//如果类型true正确,v返回的类型值本身
	fmt.Println(v, ok) //hello true
	v1, ok1 := x1.([]int)
	//如果类型false错误,v返回的空对象
	fmt.Println(v1, ok1) //[] false

}
