package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

//定义存储header的结构体

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func main() {

	//声明头部信息
	headerData := map[string]string{
		"typ": "JWT",
		"alg": "HS256",
	}

	//对头部信息序列化
	headerByte, _ := json.Marshal(headerData)

	//base64编码
	headerString := strings.TrimRight(base64.URLEncoding.EncodeToString(headerByte), "=")

	//打印base64编码
	fmt.Println(headerString) //eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9

	//base64解码
	headerUnByte, _ := base64.URLEncoding.DecodeString(headerString)
	//打印解码的内容,可以看到解码后的数据结构
	fmt.Println(string(headerUnByte)) //eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9

	//实例化结构体对象
	header := Header{}

	//将解码后的内容反序列化
	err := json.Unmarshal(headerUnByte, &header)
	if err != nil {
		return
	}

	//打印结构体数据
	fmt.Println("header>>> ", header) //header>>>  {HS256 JWT}

}
