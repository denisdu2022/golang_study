package main

import "fmt"

//class Animal {
//	birthday timestamp
//	legs int
//	mount int
//
//	walk(){
//
//}
//
//}
//
////duck  必须显示的指定Animal ,隐含的信息:duck 也具备了legs , birthday 字段,也具备了walk 方法
//class Duck implents Animal {
//
//鸭子类型强调的是事物的外部行为,而不是内部的结构  : 借鉴了动态语言的一些特性
//}

//接口的定义

type Duck interface {
	//方法的声明
	Gaga()
	Walk()
	Swimming()
}

//接口的实现

type pskDuck struct {
}

//实现方法

func (pd *pskDuck) Gaga() {
	fmt.Println("嘎嘎")
}

func (pd *pskDuck) Walk() {
	fmt.Println("this is pskDuck walking")
}

func (pd *pskDuck) Swimming() {
	fmt.Println("this is pskDuck Swimming")
}

func main() {

	//go语言的接口: 鸭子类型,在动态语言中常见 : php ,python
	/*
		go语言中处处都是interface()  ,到处都是鸭子类型 duck typing
		当看到一只鸟,走起来像鸭子,游泳起来像鸭子,叫起来也像鸭子,那么这只鸟就是鸭子
		动词: 方法 ,具备某些方法
	*/

	var d Duck = &pskDuck{}
	d.Walk()
}
