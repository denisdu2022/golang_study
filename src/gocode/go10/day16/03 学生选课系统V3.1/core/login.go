package core

import (
	. "css/databse"
	. "css/model"
	"github.com/gin-gonic/gin"
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
		ctx.String(200, "登录失败!")
	} else {
		ctx.JSON(200, gin.H{
			"msg":      "登录成功",
			"userinfo": userinfo.Account,
		})
	}

	ctx.JSON(200, gin.H{
		"msg": "ok",
	})

}
