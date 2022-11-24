package main

import "fmt"

// 玩家信息结构体
type Player struct {
	Name        string
	HealthPoint int
	Level       int
	NowPosition []int
	Prop        []string
}

// 玩家信息结构体模拟构造函数
func NewPlayer(name string, hp int, level int, np []int, prop []string) *Player {
	return &Player{
		Name:        name,
		HealthPoint: hp,
		Level:       level,
		NowPosition: np,
		Prop:        prop,
	}
}

// 发起攻击方法
func (p *Player) attack() {
	fmt.Printf("%s发起攻击! \n", p.Name)

}

// 被攻击方法
func (p *Player) attacked() {
	fmt.Printf("%s被攻击...\n", p.Name)
	fmt.Printf("血量为: %d\n", p.HealthPoint-10)
}

// 购买道具方法
func (p *Player) buyProp(prop string) {
	fmt.Printf("%s购买了%s道具", p.Name, prop)
	p.Prop = append(p.Prop, prop)
}

func main() {
	//实例化玩家对象
	p1 := NewPlayer("西门吹雪", 100, 99, []int{35, 66}, []string{"屠龙刀", "倚天剑", "铁布衫"})
	p1.attack()
	p2 := NewPlayer("东方不败", 100, 100, []int{61, 88}, []string{"屠龙刀", "倚天剑", "锁子甲"})
	p2.attacked()
	fmt.Printf("p2的血槽值:%d ", p2.HealthPoint) //p2的血槽值还是100
	fmt.Println()
	p1.buyProp("魔法石")
	fmt.Println()
	fmt.Println(p1.Prop) //[屠龙刀 倚天剑 铁布衫 魔法石]

}
