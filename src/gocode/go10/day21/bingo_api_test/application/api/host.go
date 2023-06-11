package api

import (
	"bingotest01/application/constants"
	"bingotest01/application/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

//创建主机

//func HostCreate(ctx *gin.Context) {
//	//模拟数据
//	ctx.JSON(http.StatusOK, gin.H{
//		"msg":    constants.CodeSuccess,
//		"host01": "127.0.0.1",
//	})
//}

//添加主机类别

func HostCategoryCreate(ctx *gin.Context) {
	//调用service
	hostCateGory, err := services.CreateHostCategory(ctx)

	//判断
	if err != nil || hostCateGory.ID < 1 {
		//记录错误日志
		zap.S().Error(err.Error())
		//返回错误信息
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeCreateHostCategoryFail,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data":    hostCateGory,
	})
}

//获取所有主机列表

func HostCategoryList(ctx *gin.Context) {
	//调用业务层获取主机类别列表
	hostCategoryList, err := services.GetHostCategoryList(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeGetHostCategoryFail,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data":    hostCategoryList,
	})
}

//添加主机

func HostCreate(ctx *gin.Context) {
	//调用业务层创建主机
	host, err := services.CreateHost(ctx)
	//判断
	if err != nil || host.ID < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeCreateHostFail,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"host": host,
		},
	})
}

//查看主机

func HostList(ctx *gin.Context) {
	//接收查询参数
	name, _ := ctx.GetQuery("name")
	hostname, _ := ctx.GetQuery("host")
	hostCategoryId, _ := strconv.Atoi(ctx.Query("host_category_id"))
	//调用业务层接口,把接收的参数传入业务层接口
	hostList, err := services.GetHostList(name, uint(hostCategoryId), hostname)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeGetHostFail,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"host_list": hostList,
		},
	})
}

//删除主机

func HostDelete(ctx *gin.Context) {
	//获取参数:要删除的主机ID
	idStr, _ := ctx.GetQuery("id")
	idInt, _ := strconv.Atoi(idStr)
	//调用业务层删除方法
	host, err := services.DeleteHost(uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeDelHostFail,
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"已成功删除主机:": host.Name,
		},
	})
}
