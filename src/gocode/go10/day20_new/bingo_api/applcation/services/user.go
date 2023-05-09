package services

import (
	"bingo_api/applcation/constants"
	"bingo_api/applcation/model"
	"bingo_api/applcation/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//创建用户

func CreateUser(ctx *gin.Context) (uint, error) {
	//实例化User对象
	user := model.User{}
	//定义error对象
	var err error
	//接收前端发来的数据
	if err = ctx.ShouldBindJSON(&user); err != nil {
		return 0, err
	}

	user.HashPassword, err = utils.MakeHashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	return user.Insert()
}

//获取指定ID的用户

func GetUserById(id uint) (user model.User) {
	//实例化user对象
	user = model.User{}
	//调用指定ID获取用户方法
	user.GetOneByID(id)
	//返回user对象
	return user
}

//根据账户信息(用户名,手机,邮箱)获取账户

func GetUserByAccount(account string) (user model.User) {
	//实例化user对象
	user = model.User{}
	//调用user方法
	user.GetOneByAccount(account)
	//返回user对象
	return user
}

//获取所有用户

func GetAllUser() []model.User {
	//实例化user对象
	user := model.User{}
	return user.GetAll()
}

//更新用户密码

func ChangeUserPassword(user model.User, RawPassword string) {
	password, _ := utils.MakeHashPassword(RawPassword)
	user.ChangePassword(password)
}

//用户登录

func UserLogin(ctx *gin.Context) (user model.User, err error) {
	//实例化user对象
	user = model.User{}
	//获取user
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return user, err
	}

	//数据库查询用户名
	user.GetOneByAccount(user.Username)
	//判断:如果用户id小于1则返回用户不存在
	if user.ID < 1 {
		err = errors.New(constants.NoSuchUser)
		//记录日志
		zap.S().Error(constants.NoSuchUser)
		return
	}
	//判断:如果用户密码错误,则给用户返回密码错误
	checkPass := utils.CheckPassword(user.HashPassword, user.Password)
	fmt.Println("checkPass >>>> ", checkPass)
	if !checkPass {
		err = errors.New(constants.PasswordError)
		//记录日志
		zap.S().Error(constants.PasswordError)
		return
	}

	return
}
