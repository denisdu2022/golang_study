package main //声明文件所在的包,每个go文件必须有所在的包
import "fmt" //引入程序中需要的包,为了使用包下的函数,比如:Println

//func fool() (int, string) {
//	//return 12, "jieke"
//	x := 12
//	y := "jieke"
//
//	return x, y
//}

func main() { //main 主函数 程序的入口

	//打印hello word
	//fmt.Println("hello world !!!")

	// 这是行注释

	/*
	   这是块注释
	   块注释
	   注释
	*/

	//var name string
	//var age int
	//var isok bool
	//fmt.Println(name) //""空字符串
	//fmt.Println(age)  //0
	//fmt.Println(isok) //false

	////声明多个同类型变量
	//var x, y int
	////声明多个不同类型的变量
	//var (
	// a1 int
	// a2 string
	// a3 bool
	//)
	//
	//fmt.Println(x, y, a1, a2, a3)

	//1.先声明在赋值
	//var name string
	////var name string = "hello"   //变量不能重复声明
	//name = "hello"
	//
	//fmt.Println(name)

	//2.直接声明赋值
	//var name string = "jieke"
	//var age int = 22
	//fmt.Println

	//3.声明精简变量
	//s1 := "jieke"
	//age := 22
	//fmt.Println(s1, age)

	//4.一行赋值多个变量
	//var name, age, shengao = "jieke", 22, 178
	//fmt.Println(name, age, shengao)

	//5.变量名 = 变量名 值拷贝
	//var a = 100
	//var b = a
	//fmt.Println(a, b) //a和b都是100
	//b = 200
	//fmt.Println(b)

	//6.变量名 = 值 + 值 变量名
	//var a, b = 10, 20
	//var c = a + b //数字做运算
	//fmt.Println(c)
	//
	//var s1, s2 = "name: ", "jieke"
	//var s3 = s1 + s2 //字符串相加等于拼接
	//fmt.Println(s3)

	//lianxi:将x，y两个变量的值进行交换
	//y, _ := fool()
	//_, x := fool()
	//fmt.Println(x)
	//fmt.Println(y)

	//常量
	//const pi = 3.1415926
	//const e = 2.17
	//声明多个常量
	//const (
	//	pi = 3.14
	//	a  = 23
	//	b  = 33
	//)

	//const (
	//	n1 = 100
	//	b1
	//	c1
	//)
	//fmt.Println(n1, b1, c1)

	//var a, b = 10, 20
	//a, b = b, a //有局限只能是相同类型
	////fmt.Println(a, b)
	//const (
	//	n1 = iota //当iota出现的时候为0
	//	n2        //1
	//	n3        //2
	//	n4        //3
	//)
	//fmt.Println(n1, n2, n3, n4)

	//	const (
	////		n1 = iota //0
	////		n2        //1
	////		_
	////		n4 //3
	////	)
	////
	////	fmt.Println(n1, n2, n4)
	////}

	//const (
	//	n1 = iota //0
	//	n2 = 100  //100
	//	n3 = iota //2
	//	n4        //3
	//)
	//fmt.Println(n1, n2, n3, n4)

	//const (
	//	n1 = iota    //0
	//	n2 = 1 << 10 //这里的<<表示左移操作,1<<10表示将1的二进制位向左移10位,也就是由1变成了100000000000
	//	//也就是十进制的1024
	//	n3 //1024
	//	n4 //1024,声明多个变量的时候不赋值,默认是上一个的值
	//	n5 = 200
	//	n6
	//	n7 = 2 << 2 //表示将二进制向左移2位也就是10,变成了1000,也就是10进制的8
	//)
	//fmt.Println(n1, n2, n3, n4, n5, n6, n7)

	//const (
	//	_  = iota
	//	KB = 1 << (10 * iota)
	//	MB = 1 << (10 * iota)
	//	GB = 1 << (10 * iota)
	//	TB = 1 << (10 * iota)
	//	PB = 1 << (10 * iota)
	//)
	//fmt.Println("KB :", KB)
	//fmt.Println("MB :", MB)
	//fmt.Println("GB :", GB)
	//fmt.Println("TB :", TB)
	//fmt.Println("PB :", PB)

	//const (
	//	//当iota出现是为0
	//	a, b = iota + 1, iota + 2 //1,2  iota为0+1 = 1  iota 为0 +2 =2
	//	c, d                      //2,3  第二行,上一行iota已经是1了,所以c是2 ,d是iota 的当前值为+1
	//	e, f                      //3,4 第三行,上一行iota已经是2了,所以e是3 ,f是iota 的当前值为+1
	//)
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
	//fmt.Println(e)
	//fmt.Println(f)

	var a = 10
	var b = 20
	var t = a
	a = b
	b = t

	fmt.Println(a, b)

}
