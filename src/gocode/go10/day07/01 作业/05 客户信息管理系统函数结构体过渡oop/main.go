package main

import "fmt"

type Person struct {
	name   string
	dengji int8
}

type Hero struct {
	Person
	healthy int8
	attack  string
	weapon  Weapon
}

type Weapon struct {
	name   string
	color  string
	attack int
}

func (w Weapon) shejian() {
	fmt.Println("正在射箭🏹...")
}

func main() {

	//实例化玩家对象
	//把一个结构体对象,赋值给一个结构体成员变量  链式操作
	h1 := Hero{healthy: 100, Person: Person{"红孩儿", 88}, attack: "发起攻击", weapon: Weapon{name: "红缨枪", color: "red", attack: 80}}
	fmt.Printf("玩家:%s 等级:%d 动作:%s 使用的武器:%s  武器颜色:%s  武器攻击力:%d  ", h1.name, h1.dengji, h1.attack, h1.weapon.name, h1.weapon.color, h1.weapon.attack)
	h1.weapon.shejian()
	fmt.Println()

}
