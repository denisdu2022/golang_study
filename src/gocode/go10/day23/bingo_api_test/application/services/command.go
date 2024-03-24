package services

import (
	"bingotest01/application/model"
	"github.com/gin-gonic/gin"
)

//添加指令模板类别

func CreateCmdCategory(ctx *gin.Context) (model.CommandCategoryInstance, error) {
	//实例化指令模板类别对象
	commandCategory := model.CommandCategory{}
	//实例化指令模板类别实例
	var instance model.CommandCategoryInstance
	var err error
	//从前端数据库获取到的指令模板类别
	if err = ctx.ShouldBindJSON(&commandCategory); err != nil {
		return instance, err
	}
	//入库
	err = commandCategory.Insert()
	//对实例化指令模板类别实例进行重新赋值
	instance = model.CommandCategoryInstance{
		ID:   commandCategory.ID,
		Name: commandCategory.Name,
	}
	//返回对实例化指令模板类别实例
	return instance, err
}

//获取指令模板类别

func GetCmdCategoryList(ctx *gin.Context) ([]model.CommandCategoryInstance, error) {
	commandCategory := model.CommandCategory{}
	//从数据库查询指令模板类别
	commandCategoryList, err := commandCategory.GetAll()
	return commandCategoryList, err
}

//添加指令模板

func CreateCmd(ctx *gin.Context) (model.CommandInstance, error) {
	//实例化指令模板
	cmd := model.Command{}
	//实例化指令模板实例
	var instance model.CommandInstance
	var err error
	//实例化指令模板类别
	var cmdCategory model.CommandCategory

	//从前端数据库获取到的指令模板
	if err := ctx.ShouldBindJSON(&cmd); err != nil {
		return instance, err
	}
	//入库
	err = cmd.Insert()
	if err != nil {
		return instance, err
	}
	//使用指令模板类别ID查询指令模板类别
	err = cmdCategory.GetOneById(cmd.CommandCategoryId)
	if err != nil {
		return instance, err
	}
	//实例化指令模板实例,重新赋值,回写
	instance = model.CommandInstance{
		ID:           cmd.ID,
		Name:         cmd.Name,
		Content:      cmd.Content,
		CategoryID:   cmdCategory.ID,
		CategoryName: cmdCategory.Name,
	}
	return instance, err
}

//获取指令模板列表

func GetCmdList(name string, cmdCategoryId uint) ([]model.CommandInstance, error) {
	//实例化命令列表
	var cmdList []model.Command
	//实例化命令列表实例
	var instanceList []model.CommandInstance
	var cmd model.Command
	var err error
	//从数据库查询实例化命令列表
	cmdList, err = cmd.GetAll(name, cmdCategoryId)
	//查询到多个之后去循环赋值
	for _, cmdItem := range cmdList {
		instanceList = append(instanceList, model.CommandInstance{
			ID:           cmdItem.ID,
			Name:         cmdItem.Name,
			Content:      cmdItem.Content,
			CategoryID:   cmdItem.CommandCategory.ID,
			CategoryName: cmdItem.CommandCategory.Name,
		})
	}
	return instanceList, err

}

//获取指令模板信息

func GetCmd(id int) (model.CommandInstance, error) {
	cmd := model.Command{}
	//查询
	err := cmd.GetOneById(uint(id))
	instance := model.CommandInstance{
		ID:           uint(id),
		Name:         cmd.Name,
		Content:      cmd.Content,
		CategoryID:   cmd.CommandCategory.ID,
		CategoryName: cmd.CommandCategory.Name,
	}
	return instance, err
}
