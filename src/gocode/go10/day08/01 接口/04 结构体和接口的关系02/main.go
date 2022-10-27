package main

import "fmt"

//洗衣机接口
type WashingMachine interface {
	wash() //洗衣服
	dry()  //甩干方法
}

//甩干器
type Dryer struct {
	brand string
}

//实现洗衣机甩干接口的dry()甩干方法
func (w Dryer) dry() {
	fmt.Printf("%s甩干衣服...\n", w.brand)
}

//海尔洗衣机
type Haier struct {
	name  string
	Dryer //嵌入甩干器
}

//实现洗衣机接口的wash方法
func (h Haier) wash() {
	fmt.Printf("%s洗衣服...\n", h.brand)
}

func main() {
	//实例化洗衣机接口对象
	var wm WashingMachine
	wm = Haier{
		name:  "海尔洗衣机",
		Dryer: Dryer{brand: "西门子"},
	}

	wm.wash() //西门子洗衣服...
	wm.dry()  //西门子甩干衣服...

}
