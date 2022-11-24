package main

import "fmt"

////Dog类型结构体
//type Dog struct {
//	Name string
//	Kind string
//}
////吃的方法
//func (d Dog)eat()  {
//	fmt.Printf("%s正在吃,品种:%s \n",d.Name,d.Kind)
//}
////睡得方法
//func (d Dog)sleep()  {
//	fmt.Printf("%s正在睡,品种:%s \n",d.Name,d.Kind)
//
//}//Cat类型结构体
//type Cat struct {
//	Name string
//	Kind string
//}
//
////吃的方法
//func (c Cat)eat()  {
//	fmt.Printf("%s正在吃,品种:%s \n",c.Name,c.Kind)
//}
////睡得方法
//func (c Cat)sleep() {
//	fmt.Printf("%s正在睡,品种:%s \n", c.Name, c.Kind)
//}

//以上代码方法重复

//--------------------------------------------------------

// 结构体的继承
// 动物类型结构体
type Animal struct {
	Name string
}

// 动物所有的方法
func (a *Animal) eat() {
	fmt.Printf("%s正在吃... \n", a.Name)
}

func (a Animal) sleep() {
	fmt.Printf("%s正在睡... \n", a.Name)
}

// Dos类型结构体
type Dog struct {
	Kind string
	Animal
}

// 属于Dog的方法
func (d *Dog) bark() {
	fmt.Printf("%s正在bark,品种:%s... \n", d.Name, d.Kind)
}

// Cat类型结构体
type Cat struct {
	Kind string
	Animal
}

func main() {
	//实例化Dog对象
	d1 := Dog{"金毛", Animal{"小白"}}
	d1.eat()   //小白正在吃...
	d1.sleep() //小白正在睡...
	d1.bark()  //小白正在bark,品种:金毛...

	//实例化Cat对象
	c1 := Cat{"狮子王", Animal{"喵喵"}}
	c1.eat()   //喵喵正在吃...
	c1.sleep() //喵喵正在睡...

}
