package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

//创建一个存储解码后的payload的结构体

type PayLoad struct {
	Sub   string `json:"sub"`
	Exp   string `json:"exp"`
	Name  string `json:"name"`
	Admin string `json:"admin"`
	Info  string `json:"info"`
}

func main() {

	//定义一个payload对象
	payloadData := map[string]string{
		"sub":   "1234567890",
		"exp":   "0987654321",
		"name":  "masterChief",
		"admin": "1",
		"info":  "abc12232",
	}

	//将payload信息序列化
	payloadByte, _ := json.Marshal(payloadData)

	//进行base64编码
	payloadBase64 := strings.TrimRight(base64.URLEncoding.EncodeToString(payloadByte), "=")
	//打印编码后的内容
	//eyJhZG1pbiI6IjEiLCJleHAiOiIwOTg3NjU0MzIxIiwiaW5mbyI6ImFiYzEyMjMyIiwibmFtZSI6InNoaWd1YW56aGFuZyIsInN1YiI6IjEyMzQ1Njc4OTAifQ
	fmt.Println("payloadBase64>>> ", payloadBase64)

	//base64解码
	payloadUnBase64, _ := base64.URLEncoding.DecodeString(payloadBase64)
	//打印解码后的数据
	fmt.Println("打印解码后的数据", string(payloadUnBase64))

	//实例化结构体对象
	payload := PayLoad{}

	//反序列化
	err := json.Unmarshal(payloadUnBase64, &payload)
	if err != nil {
		return
	}
	//打印已解码存储的数据
	fmt.Printf("payload>>> %#v\n", payload)
	//可以获取对象里的内容
	fmt.Println("存储的对象的名字: ", payload.Name)
}
