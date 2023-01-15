package core

import "github.com/gin-gonic/gin"

func GetClass(ctx *gin.Context) {
	ctx.HTML(200, "class", nil)
}
