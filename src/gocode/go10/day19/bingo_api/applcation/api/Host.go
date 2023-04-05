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
