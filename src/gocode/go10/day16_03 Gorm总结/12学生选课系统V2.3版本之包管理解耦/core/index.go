package core

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//首页路径跳转

func Index(ctx *gin.Context) {
	//响应首页页面
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// 首页路由函数

func GetIndex(ctx *gin.Context) {
	//2.判断之前是否登录过 cookie验证
	//isLogin, _ := ctx.Cookie("isLogin")
	//if isLogin == "true" {
	//	ctx.Redirect(http.StatusMovedPermanently, "/index")
	//} else {
	//	//响应
	//	ctx.Redirect(http.StatusMovedPermanently, "/login")
	//}

	//3.session验证
	session := sessions.Default(ctx)
	//获取session
	userID := session.Get("user_id")
	fmt.Println("user_id>>> ", userID)
	//判断
	if userID != nil {
		ctx.Redirect(http.StatusMovedPermanently, "/index")
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	}

}
