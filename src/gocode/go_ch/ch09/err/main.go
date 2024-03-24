package main

import "fmt"

func mPrint(datas ...interface{}) {
	for _, value := range datas {
		fmt.Println(value)
	}
}

func mPrint2(data interface{}) {
	fmt.Println(data)
}

func main() {
	//var data = []interface{}{
	//	"libai", 18, 1.80,
	//}
	//mPrint(data...)

	var data1 = []string{
		"data1", "data2", "data3",
	}
	//mPrint(data1 ...) //不可以的
	//mPrint2(data1)

	var datai []interface{}
	for _, value := range data1 {
		datai = append(datai, value)
	}
	fmt.Println(datai)
	mPrint(datai...)

	//var err error //error的本质就是一个interface
}
