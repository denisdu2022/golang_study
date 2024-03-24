package main

import (
	"bingotest01/application/utils"
	"fmt"
)

func main() {

	Host := "61.171.106.235"
	User := "root"
	Password := "CIAy2k0308@"
	Port := 22
	Key := ""
	Mode := "password"

	//首次使用SSH连接使用密码
	cli := utils.NewSSH(Host, User, Password, Key, Mode, Port)
	fmt.Printf("主机密码: %s , 用户名: %s ,主机IP: %s\n", cli.Password, cli.Username, cli.IP)
	//ssh连接
	if err := cli.Connect(); err != nil {
		println("主机连接失败!,错误信息是: ", err.Error())
		return
	} else {
		println("主机连接成功!")
	}

	//执行远程命令
	output, err := cli.Run("pwd")
	if err != nil {
		println("命令执行失败!")
	}

	fmt.Printf("pwd: %v\n", output)

}
