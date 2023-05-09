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
	headerData := map[string]string{
		"typ": "JWT",
		"alg": "HS256",
	}
	//header序列化
	headerByte, _ := json.Marshal(headerData)
	//header Base64编码
	headerBase64 := base64.URLEncoding.EncodeToString(headerByte)
	//打印编码后的内容
	fmt.Println(headerBase64) //eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9

	//payload
	payloadData := map[string]string{
		"sub":   "123456789",
		"exp":   "12332145",
		"name":  "masterChief",
		"admin": "1",
		"info":  "abc12321",
	}
	//payload序列化
	payloadByte, _ := json.Marshal(payloadData)
	//payload Base64编码
	payloadBase64 := base64.URLEncoding.EncodeToString(payloadByte)
	//打印编码后的内容 //eyJhZG1pbiI6IjEiLCJleHAiOiIxMjMzMjE0NSIsImluZm8iOiJhYmMxMjMyMSIsIm5hbWUiOiJtYXN0ZXJDaGllZiIsInN1YiI6IjEyMzQ1Njc4OSJ9
	fmt.Println(payloadBase64)

	//signature 签证
	//创建secret随机字符串
	secret := "secret"
	//转换类型
	key := []byte(secret)
	//加盐,对key加密
	h := hmac.New(sha256.New, key)
	//将header和payload以.拼接
	sign := fmt.Sprintf("%s.%s", headerBase64, payloadBase64)
	//写入 sign
	h.Write([]byte(sign))
	//取出signature     hex 16进制   h.sum是和
	signature := hex.EncodeToString(h.Sum(nil))
	//打印signature
	fmt.Printf("signature>>> %#v", signature) //signature>>> "122d17f44021b9d9ad57a503f05be4e123f24da395ec101d47f1cb999b16dbaf"

	//JWT
	jwt := fmt.Sprintf("%s.%s.%s", headerBase64, payloadBase64, signature)
	fmt.Println()
	//打印jwt
	fmt.Println("jwt>>> ", jwt)
}
