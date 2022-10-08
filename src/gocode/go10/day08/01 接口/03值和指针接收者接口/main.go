package main

import "fmt"

//Animal接口
type Animal interface {
	sleep() //睡得方法
}

//Dog类
type Dog struct {
	name string
}

//Dog的方法接收器
func (d Dog) sleep() {
	fmt.Printf("%s正在侧着睡...\n", d.name)
}

//Cat类
type Cat struct {
	name string
}

//Cat的方法接收器
func (c *Cat) sleep() {
	fmt.Printf("%s正在趴着睡....\n", c.name)
}

func main() {

	//实例化Animal对象
	//接收器方法参数是结构体对象
	var a Animal
	var alex = Dog{"alex"}
	a = alex         //值拷贝
	a.sleep()        //alex正在侧着睡...
	a = &Dog{"rain"} //将地址赋值给a
	a.sleep()        //rain正在侧着睡...

	fmt.Println("-------------------------------------------------------")
	//接收器方法参数是结构体对象指针
	var a2 Animal

	//a2 = Cat{"alvin"}
	//a2.sleep()
	//这样是报错的

	a2 = &Cat{"alvin"}
	a2.sleep() //alvin正在趴着睡....

}
