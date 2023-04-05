package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserAuthenticate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"func": "UserAuthenticate",
	})
}
