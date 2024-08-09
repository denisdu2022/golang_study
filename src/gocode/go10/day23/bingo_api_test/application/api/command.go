package api

import (
	"bingotest01/application/constants"
	"bingotest01/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加指令模板类别

func CmdCategoryCreate(ctx *gin.Context) {
	//调用业务层完成指令模板类别的添加操作
	cmdCategory, err := services.CreateCmdCategory(ctx)
	if err != nil || cmdCategory.ID < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeCreateCmdCategoryFail,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"cmdCategory": cmdCategory,
		},
	})

}

//获取所有指令模板类别

func CmdCategoryList(ctx *gin.Context) {
	//调用业务层获取指令模板类别列表
	cmdCategoryList, err := services.GetCmdCategoryList(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeCreateCmdCategoryFail,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"cmd_category_list": cmdCategoryList,
		},
	})
}

//添加指令模板

func CmdCreate(ctx *gin.Context) {
	//调用业务层创建指令模板信息
	cmd, err := services.CreateCmd(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeCreateCmdFail,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"cmd": cmd,
		},
	})
}

//指令模板列表

func CmdList(ctx *gin.Context) {
	//接收查询指令模板列表时的筛选参数
	name, _ := ctx.GetQuery("name")
	cmdCategoryId, _ := strconv.Atoi(ctx.Query("command_category_id"))
	//调用业务层获取指令模板列表
	cmdList, err := services.GetCmdList(name, uint(cmdCategoryId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeGetCmdFail,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"cmd_list": cmdList,
		},
	})
}

//指令模板信息

func CmdInfo(ctx *gin.Context) {

	//根据指令模板ID获取指令模板信息
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeGetCmdFail,
			"message": err.Error(),
		})
		return
	}
	//根据模板ID获取指令
	cmd, err := services.GetCmd(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeGetCmdFail,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"cmd": cmd,
		},
	})

}
