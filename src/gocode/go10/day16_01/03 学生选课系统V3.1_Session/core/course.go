package core

import (
	. "css/databse"
	. "css/model"
	"github.com/gin-gonic/gin"
)

func GetCourse(ctx *gin.Context) {
	//查询course表中的记录
	var course []Course
	DB.Find(&course)
	ctx.HTML(200, "course", gin.H{
		"course": course,
	})
}
