package main //声明文件所在的包,每个go文件必须有所在的包

import "fmt" //引入程序中需要的包,为了使用包下的函数,比如:println

func main() { //main 主函数  程序的入口
	fmt.Println("HELLO GO !!!") //在控制台打印输出的字符串
	fmt.Println("皮卡皮卡丘")
	fmt.Println("皮卡皮卡丘")
	fmt.Println("皮卡皮卡丘")
	fmt.Println("皮卡皮卡丘")
	fmt.Println("皮卡皮卡丘")
	var add = 22 + 3
	fmt.Println(add)
}

// 这是行注释

/*

aaaa

fdfd
sdsds

这是块注释(多行注释)





*/
