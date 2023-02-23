package core

import (
	. "cssv2.3/database"
	. "cssv2.3/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// login页面

func GetLoginHtml(ctx *gin.Context) {
	//响应
	ctx.HTML(http.StatusOK, "login.html", nil)
}

//login登录

func PostLogin(ctx *gin.Context) {
	//获取前端的用户名密码
	name := ctx.PostForm("userName")
	pwd := ctx.PostForm("passWord")
	fmt.Println(name, pwd)
	//数据库校验
	var userInfo UserInfo
	DB.Where("account = ? and pwd = ?", name, pwd).Take(&userInfo)

	if userInfo.ID != 0 {
		//01.设置cookie
		ctx.SetCookie("isLogin", "true", 2000000, "/", "127.0.0.1", false, true)
		//响应
		ctx.Redirect(http.StatusMovedPermanently, "/index")
	} else {
		//响应
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	}

}
