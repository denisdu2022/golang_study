package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//首页路径跳转

func Index(ctx *gin.Context) {
	//响应
	ctx.Redirect(http.StatusMovedPermanently, "/login")
}

// 首页路由函数

func GetIndex(ctx *gin.Context) {
	//2.判断之前是否登录过
	isLogin, _ := ctx.Cookie("isLogin")
	if isLogin == "true" {
		//响应首页页面
		ctx.HTML(http.StatusOK, "index.html", nil)
	} else {
		//响应
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	}

}
