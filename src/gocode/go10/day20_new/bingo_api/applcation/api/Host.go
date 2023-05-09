package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建主机

func HostCreate(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"func": "HostCreate",
	})
}

//查看主机列表

func HostList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"host01": "127.0.0.3",
		"host02": "127.0.0.4",
	})
}
