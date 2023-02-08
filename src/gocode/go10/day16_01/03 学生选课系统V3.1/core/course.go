package core

import "github.com/gin-gonic/gin"

func GetCourse(ctx *gin.Context) {
	ctx.HTML(200, "course", nil)
}
