package middleware

import (
	"github.com/gin-contrib/cors"
	"time"
)

var Cors = cors.New(cors.Config{
	//准许跨域请求网站,多个使用,分开,可以使用*通配符
	AllowOrigins: []string{"*"},
	//准许使用的请求方式
	AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
	//准许使用的请求头
	AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
	//显示的请求表头
	ExposeHeaders: []string{"Content-Type"},
	//凭证共享,确定共享
	AllowCredentials: true,
	//容允跨域的原点网站,可以直接return true
	AllowOriginFunc: func(origin string) bool {
		return true
	},
	//超时时间设定
	MaxAge: 24 * time.Hour,
})
