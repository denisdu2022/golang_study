package api

import (
	"bingotest01/application/constants"
	"bingotest01/application/model"
	"bingotest01/application/services"
	"bingotest01/application/utils"
	"fmt"
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

//主机web ssh终端

func Console(ctx *gin.Context) {
	fmt.Println("----------------------------------------------------")
	//升级get请求为websocket协议
	ws, err := utils.UpGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeFail,
			"message": err.Error(),
		})
		return
	}

	//根据ID获取主机连接信息
	id, err := strconv.Atoi(ctx.Param("id"))
	fmt.Println("主机id>>> ", id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeFail,
			"message": err.Error(),
		})
		return
	}

	//实例化主机对象
	host := model.Host{}
	//根据主机ID去数据库查询主机
	err = host.GetOneById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeFail,
			"message": err.Error(),
		})
		return
	}
	//主机信息
	fmt.Println("host>>> ", host)

	//创建SSH远程连接
	cli := utils.NewSSH(host.IpAddr, host.Username, "", host.PrivateKey, "key", int(host.Port))
	if err := cli.Connect(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeFail,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("秘钥登录成功!!!")
	fmt.Println("----------------------------------------------------")
	cli.WS2SSH(ws)
}

//远程主机批量执行命令

func HostRunCmdList(ctx *gin.Context) {
	//实例化param
	var param = model.CmdParam{}
	//从上下文对象中获取前端传过来的主机信息
	err := ctx.ShouldBindJSON(&param)
	fmt.Println("param>>> ", param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.ParamError,
			"message": err.Error(),
		})
		return
	}
	//调用service函数
	err = services.ExecuteCMDList(param.HostIds, param.Cmd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.ExecuteCMDFail,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
	})
}
