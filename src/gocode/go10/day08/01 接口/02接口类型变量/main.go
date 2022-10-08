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
func (c Cat) sleep() {
	fmt.Printf("%s正在趴着睡....\n", c.name)
}

func foo(a Animal) {
	a.sleep()
}

func main() {

	//实例化Animal对象
	var a Animal
	var alex = Dog{"alex"}
	a = alex  //值拷贝
	a.sleep() //alex正在侧着睡...

	var rain = Cat{"rain"}
	a = rain
	a.sleep() //rain正在趴着睡....
	fmt.Println("---------------------------")
	foo(alex) //alex正在侧着睡...
	foo(rain) //rain正在趴着睡....
}
