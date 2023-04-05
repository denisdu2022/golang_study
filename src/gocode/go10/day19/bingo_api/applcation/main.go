package main

import "bingo_api/applcation/initialize"

func main() {

	//1.初始化路由
	Router := initialize.InitRouter()

	//2.启动,默认监听8080端口
	Router.Run()

}
