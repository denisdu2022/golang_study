package main

import "api_01/application/initialize"

func main() {

	//路由初始化
	r := initialize.InitRouter()

	//启动 端口8081
	r.Run(":8081")

}
