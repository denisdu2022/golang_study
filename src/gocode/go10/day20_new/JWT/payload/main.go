package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

//定义payload结构体

type Payload struct {
	Sub   string `json:"sub"`
	Exp   string `json:"exp"`
	Name  string `json:"name"`
	Admin string `json:"admin"`
	Info  string `json:"info"`
}

func main() {

	//定义一个payload的map对象
	payloadData := map[string]string{
		"sub":   "123456789",
		"exp":   "12332145",
		"name":  "masterChief",
		"admin": "1",
		"info":  "abc12321",
	}
	//将payload序列化
	payloadByte, _ := json.Marshal(payloadData)

	//base64编码
	payloadBase64 := strings.TrimRight(base64.URLEncoding.EncodeToString(payloadByte), "=")
	//打印编码后的内容
	fmt.Println(payloadBase64) //eyJhZG1pbiI6IjEiLCJleHAiOiIxMjMzMjE0NSIsImluZm8iOiJhYmMxMjMyMSIsIm5hbWUiOiJtYXN0ZXJDaGllZiIsInN1YiI6IjEyMzQ1Njc4OSJ9

	//base64解码
	payloadUnBase64, _ := base64.URLEncoding.DecodeString(payloadBase64)
	//实例化结构体对象
	payload := Payload{}
	//对解码后的数据反序列化
	json.Unmarshal(payloadUnBase64, &payload)

	//打印payload
	fmt.Printf("payload>> %#v", payload) //payload>> main.Payload{Sub:"123456789", Exp:"12332145", Name:"masterChief", Admin:"1", Info:"abc12321"}

}
