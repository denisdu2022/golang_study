package main

import (
	"bingotest01/application/utils"
	"fmt"
)

func main() {

	Host := "61.171.106.235"
	User := "root"
	Password := ""
	Port := 22
	Mode := "key"

	//key里放私钥
	Key := "-----BEGIN OPENSSH PRIVATE KEY-----\n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\n-----END OPENSSH PRIVATE KEY-----"

	cli := utils.NewSSH(Host, User, Password, Key, Mode, Port)

	//免密远程连接
	if err := cli.Connect(); err != nil {
		println("远程连接失败,失败原因是: ", err)
		return
	} else {
		println("远程连接成功!私钥认证成功!")
	}

	defer cli.Client.Close()

	//执行命令
	output, err := cli.Run("pwd")
	if err != nil {
		println("命令执行失败!")
		return
	}
	fmt.Println("命令执行成功!,output: ", output)
}
