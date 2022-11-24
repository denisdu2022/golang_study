package main

import "fmt"

// Animal接口
type Animal interface {
	Run
	Jump
}

// 跑的接口
type Run interface {
	run() string
}

// 跳的接口
type Jump interface {
	jump() string
}

// 人这个类
type Person struct {
	name string
}

// 汽车这个类
type Car struct {
	brand string
}

// 汽车实现跑的方法接收器
func (c Car) run() string {
	fmt.Printf("%s汽车正在飞速的行驶...\n", c.brand)
	return "跑"
}

// 跑的方法接收器
func (p Person) run() string {
	fmt.Printf("%s正在跑...\n", p.name)
	return "Run"
}

// 跳的方法接收器
func (p Person) jump() string {
	fmt.Printf("%s正在跳...\n", p.name)
	return "Jump"
}

func main() {

	//实例化人这个类的对象
	var daerwen = Person{"daerwen"}

	var r Run
	r = daerwen
	r.run() //是可以的
	//r.jump //是不可以的因为r里边不包含jump的方法

	//就是可以的,因为Person一个结构体里包含了run和jump方法
	daerwen.jump() //正在跳...
	daerwen.run()  //正在跑...

	fmt.Println("---------------------------------------------------------")
	var byd = Car{"BYD"}
	byd.run() //BYD汽车正在飞速的行驶...
	ret := byd.run()
	fmt.Println(ret) //跑
	fmt.Println("---------------------------------------------------------")
	var a Animal
	var b = Person{"bijiasuo"}
	a = b
	a.run()  //bijiasuo正在跑...
	a.jump() //bijiasuo正在跳...

}
