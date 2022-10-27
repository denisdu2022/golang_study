package main

import "fmt"

//Sleep接口
type Sleep interface {
	sleep()
}

//Runner 接口
type Runner interface {
	run()
}

//Animal接口  接口嵌套包含Sleep接口和Runner接口
type Animal interface {
	Sleep
	Runner
}

//Dog类
type Dog struct {
	name string
}

//Dog实现了run方法和sleep方法,也实现了Animal接口
func (d Dog) run() {
	fmt.Printf("%s正在吐舌头跑...\n", d.name)
}

func (d Dog) sleep() {
	fmt.Printf("%s正在仰天大睡...\n", d.name)
}

func main() {

	var a Animal
	var d = Dog{"拉布拉多"}
	a = d
	a.run()   //拉布拉多正在吐舌头跑...
	a.sleep() //拉布拉多正在仰天大睡...

}
