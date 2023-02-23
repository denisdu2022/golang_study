package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 班级路由函数

func GetClass(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "class.html", nil)
}
