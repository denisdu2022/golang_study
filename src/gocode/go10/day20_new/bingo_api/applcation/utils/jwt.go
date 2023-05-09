package utils

import (
	"bingo_api/applcation/config"
	"bingo_api/applcation/constants"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//JWT

type JWT struct {
	SigningKey []byte
}

//新建一个JWT的实例

func NewJWT() *JWT {
	return &JWT{
		[]byte(config.Conf.SecretKey),
	}
}

//payload(载荷)

type PublicClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	jwt.StandardClaims
}

//生成一个AccessToken

func (j *JWT) AccessToken(claims PublicClaims) (string, error) {
	//获取当前时间
	now := time.Now()
	//获取当前时间的Unix时间戳
	nowAt := now.Unix()
	//设置token的签发时间
	claims.IssuedAt = nowAt
	//设置jwt的唯一身份标识,主要用来作为一次性token,从而回避重放攻击
	claims.Id = UUID4()
	//判断:如果配置项中有设置token的过期时间
	if config.Conf.ExpiresAt > 0 {
		//则在token的载荷信息中记录,当前token的启用时间和过期时间
		expAt := now.Add(time.Duration(config.Conf.ExpiresAt) * time.Second).Unix()
		//载荷中的过期时间
		claims.ExpiresAt = expAt
		//载荷中的启用时间
		claims.NotBefore = nowAt
	}
	//判断:如果配置中有设置token的启用时间
	if config.Conf.NotBefore > 0 {
		//则在token的载荷信息中记录,当前token的启用时间
		nbfAt := now.Add(time.Duration(config.Conf.NotBefore) * time.Second).Unix()
		//启用时间
		claims.NotBefore = nbfAt
	}

	//创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//返回 jwt的token对象
	return token.SignedString(j.SigningKey)
}

//从token中提取载荷

//func (j *JWT) GetPayloadByToken(tokenString string) (*PublicClaims, error) {
//	//解析token
//	token, err := jwt.ParseWithClaims(tokenString, &PublicClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return j.SigningKey, nil
//	})
//	//判断错误
//	if err != nil {
//		if ve, ok := err.(*jwt.ValidationError); ok {
//			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
//				return nil, errors.New(constants.TokenIsMalformd)
//			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
//				return nil, errors.New(constants.TokenIsExpired)
//			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
//				return nil, errors.New(constants.TokenIsNotValidYet)
//			} else {
//				return nil, errors.New(constants.TokenIsInvalid)
//			}
//		}
//	}
//	//判断claims
//	if claims, ok := token.Claims.(*PublicClaims); ok && token.Valid {
//		return claims, nil
//	}
//
//	return nil, errors.New(constants.TokenIsInvalid)
//}

func (j *JWT) GetPayloadByToken(tokenString string) (*PublicClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &PublicClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(constants.TokenIsMalformd)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(constants.TokenIsExpired)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(constants.TokenIsNotValidYet)
			} else {
				return nil, errors.New(constants.TokenIsInvalid)
			}
		}
	}
	if claims, ok := token.Claims.(*PublicClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New(constants.TokenIsInvalid)
}
