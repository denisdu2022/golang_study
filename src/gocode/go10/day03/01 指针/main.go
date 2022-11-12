package main

func main() {

	////00 指针的取值与取址
	//var a = 100
	//var s = "string"
	//
	////取址符: & 取值符: *
	//fmt.Println("a的内存地址是,使用取址符 /&: ", &a) //0x1400001c098
	//fmt.Println("s的内存地址是,使用取址符 /&: ", &s) //0x14000010230
	//
	//var b *int
	//b = &a
	//fmt.Println("b的值: ", b) //0x1400001c098  存的是a的内存地址
	//fmt.Println("b的内存地址是: ", &b)
	//fmt.Println("b地址对应的值是: ", *b) //100
	//
	////01 地址的格式化打印
	//var x = 10
	//fmt.Printf("%p\n", &x) //0x1400012c008
	//x = 100
	//fmt.Printf("%p\n", &x) //0x1400012c008   重新赋值后值变,内存地址不变
	//
	////取值
	//fmt.Println(*&x) //100
	//
	////02 指针的应用
	////使用等号将一个变量的值赋值给另一个变量时,如x = y,实际上是在内存中将x的值进行了拷贝,值拷贝
	//var x = 22
	//var y = x
	//var z = &x
	//x = 20
	//fmt.Println(y)  //22
	//fmt.Println(*z) //20
	//*z = 30
	//fmt.Println(x) //30
	//
	////练习 01
	//var x = 10
	//fmt.Println(&x)
	//var y = &x
	//fmt.Println(&y, *y)
	//var z = *y
	//fmt.Println(&y, *y)
	//x = 20
	//fmt.Println(&y, *y)
	//fmt.Println(x)  //20
	//fmt.Println(*y) //20
	//fmt.Println(z)  //10
	//
	////练习 02
	//var a = 100
	//var b = &a
	//var c = &b
	//
	//**c = 200
	//fmt.Println(a)  //a 200
	//fmt.Println(*b) //b 200
	//fmt.Println(*c) //c 0x1400012c008

}
