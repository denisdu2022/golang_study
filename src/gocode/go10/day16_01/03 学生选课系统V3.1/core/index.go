package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 首页

func Index(ctx *gin.Context) {
	//2. 获取cookie 判断之前是否登录成功过
	isLogin, _ := ctx.Cookie("isLogin")

	//3. 判断校验
	if isLogin == "true" {
		ctx.HTML(200, "index.html", nil)
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	}

}
