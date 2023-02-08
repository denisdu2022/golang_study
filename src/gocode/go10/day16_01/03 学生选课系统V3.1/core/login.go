package core

import (
	. "css/databse"
	. "css/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//登录

func GetLoginHtml(ctx *gin.Context) {

	ctx.HTML(200, "login.html", nil)
}

func Login(ctx *gin.Context) {

	//获取登录的用户名密码
	userName := ctx.PostForm("userName")
	passWord := ctx.PostForm("passWord")
	//校验
	//实例化UserInfo对象
	var userinfo UserInfo
	//数据库查询
	DB.Where("account = ? and password = ?", userName, passWord).Take(&userinfo)
	//判断
	if userinfo.ID == 0 {
		//登录失败重定向到登录页面
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	} else {
		//ctx.JSON(200, gin.H{
		//	"msg":      "登录成功",
		//	"userinfo": userinfo.Account,
		//})
		//1. 设置cookie  注意:要在特殊事件上设置cookie键值对
		ctx.SetCookie("isLogin", "true", 200000, "/", "127.0.0.1", false, true)

		//登录成功重定向到首页
		ctx.Redirect(http.StatusMovedPermanently, "/")
	}

	//ctx.JSON(200, gin.H{
	//	"msg": "ok",
	//})

}
