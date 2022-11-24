package main //声明包名

import (
	"fmt"
	. "mysiteinit/api"
) //导入包

func main() {
	RestfulAPI()
	fmt.Println("-----------------------------")
	RpcAPI()

}
