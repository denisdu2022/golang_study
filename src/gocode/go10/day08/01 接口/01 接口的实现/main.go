package main

import "fmt"

//Animal接口
type Animale interface {
	sleep()
}

//定义Dog和Cat两个结构体

type Dog struct {
}

type Cat struct {
}

//Dog实现Animal接口
func (d Dog) sleep() {
	fmt.Println("狗狗正在侧着睡觉...")

}

//Cat实现Animal接口

func (c Cat) sleep() {
	fmt.Println("喵喵正在趴着睡觉...")

}

func foo(animal Animale) {
	animal.sleep()
}

func main() {

	//实例化Animal对象
	var a Animale
	//Dog
	a = Dog{}
	a.sleep() //狗狗正在侧着睡觉...
	//Cat
	a = Cat{}
	a.sleep() //喵喵正在趴着睡觉...

}
