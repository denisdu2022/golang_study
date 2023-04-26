package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

//创建一个结构体,将解码后的数据存储

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

	//将声明的头部信息序列化
	headerJsonByte, _ := json.Marshal(headerData)

	//进行base64编码  strings.TrimRight是去除右边的空格
	headerString := strings.TrimRight(base64.URLEncoding.EncodeToString(headerJsonByte), "=")
	//打印base编码 //eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
	fmt.Println(headerString)

	//base64解码
	headerByte, _ := base64.URLEncoding.DecodeString(headerString)
	//打印解码的内容 可以查看到数据的结构
	fmt.Println(string(headerByte)) //{"alg":"HS256","typ":"JWT"}

	//实例化结构体对象
	header := Header{}
	//将解码的信息反序列化
	//json.Unmarshal(headerByte, &header)
	err := json.Unmarshal(headerByte, &header)
	if err != nil {
		return
	}
	//打印存储的解码后的数据
	fmt.Printf("header>>> %#v", header) //header>>> main.Header{Alg:"HS256", Typ:"JWT"}

}
