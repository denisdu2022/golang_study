package main

func main() {
	//&变量 :获取变量的地址
	//var x = 10
	//fmt.Printf("赋值之前x对应的内存地址: %p\n", &x)  //0xc0000160a8
	//x = 100
	////整形变量x的内存地址并未发生变化,只是内存地址中的值发生了变化
	//fmt.Printf("赋值之后x的内存地址: %p\n", &x)  //0xc0000160a8
	//
	////1. 获取地址: &变量
	//var x = 10 //x称之为整形变量
	//fmt.Println(&x)
	////2. 地址赋值
	//var p *int     //p 是一个指针变量,是一个整形指针类型
	//p = &x         //把x的内存地址存到整形指针变量p中
	////以上两行可以这样写
	////var p = &x
	//fmt.Println(p) //p里边存的地址值是x的内存地址
	//
	////3. 取值操作: *指针变量
	////打印存的内存地址的值,也就是说存的是x的内存地址,需要打印x内存地址里的值,打印后这个值的类型是原来的类型
	//fmt.Println(*p, reflect.TypeOf(*p))  //10 int
	////修改存的地址里的值
	//*p = 100
	//fmt.Println(x)
	//
	//思考题
	//var a = 10
	//var b = a //值拷贝
	//b = 100
	////修改b的值a并不会改变,因为a和b对应的是两块独立的空间
	//fmt.Println(a, b)
	//
	////要想实现修改b,a也会改变需要这样做
	//var a = 10
	//var b = &a      //相当与: var b *int  b = &a
	//fmt.Println(&a) //0xc0000160a8
	//var c = b       //相当与: var c *int  c = &a  通俗说:引用拷贝 ,其实也是值拷贝
	////把b里边存的a的地址值拷贝一份给c
	////c里边存的a的地址值也是指向a的内存地址
	////指针变量是引用类型
	//
	//*b = 100
	//
	//fmt.Println(a, *b)                  //100 /100
	//fmt.Println(c)                      //0xc0000160a8
	//fmt.Println(*c, reflect.TypeOf(*c)) //100 int
	//
	//指针应用1
	//var x = 10
	//var y = x //10
	//var z = &x
	//x = 20
	//fmt.Println(x)  //20
	//fmt.Println(y)  //10
	//fmt.Println(*z) //20
	//*z = 30
	//fmt.Println(x) //30
	//
	////指针应用2
	//var x = 10
	//var y = &x //把x的内存地址赋值给了y ,y是一个指针变量 *int
	//var z = *y //取y里边存的a的内存地址里的值取出来赋值给了z,所以z是10,是整形
	////z也开辟了一块内存空间里边存了10,值拷贝,x的值改变,z的值是不改变的
	//x = 20
	//fmt.Println(x)  //20
	//fmt.Println(*y) //20
	//fmt.Println(z)  //10
	//*y = 30
	//fmt.Println(z) //10
	//
	////指针应用3
	//var a = 100                    // a int 100
	//var b = &a                     //把a的内存地址赋值给b  b *int
	//var c = &b                     //把b的内存地址赋值给c  c **int
	//fmt.Println(reflect.TypeOf(c)) //**int
	//**c = 200                      // *c b的内存地址值,**c a的内存地址值 所以a 重新赋值为200
	//fmt.Println(a)                 //200
	//fmt.Println(b)                 //200
	//fmt.Println(c)                 //200
	//
	////作业
	//p1 := 1          //int类型
	//p2 := &p1        //*int
	//*p2++            //取值 自加1
	//fmt.Println(p1)  //2
	//fmt.Println(*p2) //2

}
