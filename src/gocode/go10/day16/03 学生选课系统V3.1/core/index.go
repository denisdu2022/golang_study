package core

import "github.com/gin-gonic/gin"

// 首页
func Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}
