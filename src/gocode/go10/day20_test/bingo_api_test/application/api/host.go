package api

import (
	"bingotest01/application/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建主机

func HostCreate(ctx *gin.Context) {
	//模拟数据
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    constants.CodeSuccess,
		"host01": "127.0.0.1",
	})
}
