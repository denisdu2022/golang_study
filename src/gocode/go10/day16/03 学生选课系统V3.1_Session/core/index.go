package core

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// 首页

func Index(ctx *gin.Context) {
	//cookie验证
	////2. 获取cookie 判断之前是否登录成功过
	//isLogin, _ := ctx.Cookie("isLogin")

	//3. 判断校验
	//if isLogin == "true" {
	//	ctx.HTML(200, "index.html", nil)
	//} else {
	//	ctx.Redirect(http.StatusMovedPermanently, "/login")
	//}

	//session验证
	session := sessions.Default(ctx)
	//获取session
	userId := session.Get("user_id")
	//user_id打印
	fmt.Println("user_id: ", userId, reflect.TypeOf(userId))
	//判断校验
	//if userId != 0 {
	if userId != nil {
		ctx.HTML(200, "index.html", nil)
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	}
}
