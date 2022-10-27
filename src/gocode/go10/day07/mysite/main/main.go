package main

import (
	"fmt"
	"github.com/jinzhu/now" //导入外部包
	"mysite/api"
)

func main() {
	api.RestfulMySql()
	fmt.Println("-----------------------------")
	api.Rpc()
	fmt.Println("-----------------------------")
	fmt.Println()
	fmt.Println("导入的第三方包的输出结果: ")
	fmt.Println(now.BeginningOfMinute())
}
