package main

import (
	"bingotest01/application/utils"
	"fmt"
	"io/ioutil"
)

func main() {
	Host := "61.171.106.235"
	User := "root"
	Password := "CIAy2k0308@"
	Port := 22
	Key := ""
	Mode := "password"

	//首次SSH连接使用密码
	cli := utils.NewSSH(Host, User, Password, Key, Mode, Port)

	//ssh连接
	if err := cli.Connect(); err != nil {
		println("主机连接失败,错误信息是: ", err.Error())
		return
	} else {
		println("主机连接成功!")
	}

	//向远程主机写入公钥
	rsaPubByte, err := ioutil.ReadFile("keys/id_rsa.pub")
	if err != nil {
		println("公钥读取失败")
		return
	}
	rsaPubStr := string(rsaPubByte)

	//调用方法将公钥传入远程主机
	if err := cli.AddPublicKeyToRemoteHost(rsaPubStr); err != nil {
		panic(err)
	}

	//执行命令
	res, err := cli.Run("cat ~/.ssh/authorized_keys")
	if err != nil {
		println("命令执行失败!")
		return
	}
	fmt.Println("res: ", res)
}
