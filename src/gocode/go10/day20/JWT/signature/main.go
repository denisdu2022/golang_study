package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main() {

	//header
	header := map[string]string{
		"typ": "JWT",
		"alg": "HS256",
	}
	//header序列化
	headerByte, _ := json.Marshal(header)
	//header base64编码
	headerString := base64.URLEncoding.EncodeToString(headerByte)
	//打印编码后的内容
	fmt.Println(headerString)

	//payload
	payload := map[string]string{
		"sub":   "1234567890",
		"exp":   "0987654321",
		"name":  "masterchief",
		"admin": "1",
		"info":  "123456abcdef666",
	}
	//payload序列化
	payloadByte, _ := json.Marshal(payload)
	//payload base64编码
	payloadString := base64.URLEncoding.EncodeToString(payloadByte)
	//打印编码后的内容
	fmt.Println(payloadString)

	//signature 签证
	//创建secret随机字符串
	secret := "1234567890abcdefg!@#$"
	//转换为[]byte
	key := []byte(secret)
	//加盐 对key加密
	h := hmac.New(sha256.New, key)
	//将header和payload 以.拼接
	sign := fmt.Sprintf("%s.%s", headerString, payloadString)
	h.Write([]byte(sign))
	signature := hex.EncodeToString(h.Sum(nil))
	//打印signature
	fmt.Println(signature)
	//JWT
	jwt := fmt.Sprintf("%s.%s.%s", headerString, payloadString, signature)
	//打印JWT
	fmt.Println(jwt)

}
