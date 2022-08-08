package main

import "fmt"

func main() {

	//var p *int
	////fmt.Println(p)  //<nil>
	////fmt.Println(*p)  //没有开辟空间地址
	//*p = 10 //报错

	//00 new函数
	var p = new(int)
	fmt.Println(p)  //初始的内存地址  0x14000122008
	fmt.Println(*p) //0  zero value

	*p = 10
	fmt.Println(&p, *p) //0x1400011c018 10

}
