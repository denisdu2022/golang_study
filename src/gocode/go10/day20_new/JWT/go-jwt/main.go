package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//消息提示常量

const (
	TokenExpired     = "认证令牌过期!"
	TokenNotValidYet = "认证令牌未激活!"
	TokenMalformed   = "认证令牌格式有误!"
	TokenInvalid     = "无效的认证令牌"
)

//载荷

type PublicClaims struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	//继承jwt.StandardClaims
	jwt.StandardClaims
}

//定义SecretKey秘钥

var (
	SecretKey string = "secret"
)

//获取secretKey

func GetSecretKey() string {
	return SecretKey
}

//JWT

type JWT struct {
	SigningKey []byte
}

//JWT实例对象

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSecretKey()),
	}
}

//生成AccessToken

func (j *JWT) AccessToken(claims PublicClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//从token中提取载荷

func (j *JWT) GetPayloadByToken(tokenString string) (*PublicClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &PublicClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	//判断错误
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(TokenMalformed)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(TokenExpired)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(TokenNotValidYet)
			} else {
				return nil, errors.New(TokenInvalid)
			}
		}
	}

	if claims, ok := token.Claims.(*PublicClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New(TokenInvalid)
}

//刷新token

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	//jwt.TimeFunc
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	//解析payload
	token, err := jwt.ParseWithClaims(tokenString, &PublicClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	//判断错误
	if err != nil {
		return "", err
	}
	//修改claims
	if claims, ok := token.Claims.(*PublicClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		//设置过期时间
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.AccessToken(*claims)
	}
	return "", errors.New(TokenInvalid)
}

func main() {
	//实例化载荷
	PublicClaims := PublicClaims{ID: "1", UserName: "masterChief"}
	//实例化JWT对象
	newJWT := NewJWT()

	//生成token
	token, _ := newJWT.AccessToken(PublicClaims)
	//打印生成的token
	fmt.Println(token)

	//从token中获取载荷
	payload, _ := newJWT.GetPayloadByToken(token)
	fmt.Printf("payload >>> %#v", payload)
	fmt.Println()

	//测试
	//timeNow := jwt.TimeFunc()
	//fmt.Println(timeNow)
	//fmt.Println(time.Now().Add(1 * time.Hour).Unix())

	//刷新token
	newToken, _ := newJWT.RefreshToken(token)
	//打印新的token
	fmt.Printf("NewToken: %#v", newToken)
	fmt.Println()

}
