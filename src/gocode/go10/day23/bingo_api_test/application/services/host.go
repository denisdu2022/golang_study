package services

import (
	"bingotest01/application/constants"
	"bingotest01/application/model"
	"bingotest01/application/utils"
	"bingotest01/application/validator"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

//添加主机类别

func CreateHostCategory(ctx *gin.Context) (model.HostCategoryInstance, error) {
	//实例化hostCategory对象
	hostCategory := model.HostCategory{}
	//实例化instance对象
	var instance model.HostCategoryInstance
	var err error
	//获取主机类别
	if err = ctx.ShouldBindJSON(&hostCategory); err != nil {
		return instance, err
	}
	//数据库插入主机类别
	err = hostCategory.Insert()

	instance = model.HostCategoryInstance{
		ID:   hostCategory.ID,
		Name: hostCategory.Name,
	}

	return instance, err
}

//查看主机类别

func GetHostCategoryList(ctx *gin.Context) ([]model.HostCategoryInstance, error) {
	//获取主机类别对象
	hostCategory := model.HostCategory{}
	//调用主机类别查询方法
	hostCategoryList, err := hostCategory.GetAll()
	return hostCategoryList, err
}

//添加主机

func CreateHost(ctx *gin.Context) (model.HostInstance, error) {
	//定义主机对象
	host := model.Host{}
	var err error
	//定义instance
	var instance model.HostInstance
	//定义主机类别
	var hostCategroy model.HostCategory

	//主机字段校验
	if err = validator.HostValidator(&host, ctx); err != nil {
		return instance, err
	}

	//添加的主机插入数据库
	err = host.Insert()
	if err != nil {
		return instance, err
	}

	err = hostCategroy.GetOneById(host.HostCategoryID)
	if err != nil {
		return instance, err
	}

	instance = model.HostInstance{
		ID:           host.ID,
		Name:         host.Name,
		IpAddr:       host.IpAddr,
		Port:         host.Port,
		Username:     host.Username,
		Remark:       host.Remark,
		CategoryID:   host.HostCategoryID,
		CategoryName: hostCategroy.Name,
	}
	return instance, err
}

//获取主机列表信息

func GetHostList(name string, hostCategoryId uint, hostname string) ([]model.HostInstance, error) {
	//定义主机列表
	var hostList []model.Host
	var instanceList []model.HostInstance
	var host model.Host
	var err error

	//调用查询方法
	hostList, err = host.GetAll(name, hostCategoryId, hostname)

	//range循环
	for _, hostItem := range hostList {
		instanceList = append(instanceList, model.HostInstance{
			ID:           hostItem.ID,
			Name:         hostItem.Name,
			IpAddr:       hostItem.IpAddr,
			Port:         hostItem.Port,
			Username:     hostItem.Username,
			CategoryID:   hostItem.HostCategory.ID,
			CategoryName: hostItem.HostCategory.Name,
			Remark:       hostItem.Remark,
		})
	}
	return instanceList, err
}

//删除主机

func DeleteHost(hostId uint) (model.HostInstance, error) {
	//获取要删除的host对象
	var delHost = model.Host{}
	var instance = model.HostInstance{}
	var err error

	//获取要删除的host.Id
	err = delHost.GetOneById(hostId)
	if err != nil {
		return instance, err
	}
	//调用删除方法
	err = delHost.Delete()
	if err != nil {
		return instance, err
	}

	instance = model.HostInstance{
		Name:       delHost.Name,
		IpAddr:     delHost.IpAddr,
		Port:       delHost.Port,
		Username:   delHost.Username,
		Remark:     delHost.Remark,
		CategoryID: delHost.HostCategoryID,
	}

	return instance, err
}

//执行终端命令

func ExecuteCmd(id int, cmd string) (string, error) {
	//获取主机对象
	host := model.Host{}
	//调用主机的查询方法查询主机
	err := host.GetOneById(uint(id))
	if err != nil {
		return "", err
	}

	//获取远程连接信息
	cli := utils.NewSSH(host.IpAddr, host.Username, "", host.PrivateKey, "key", int(host.Port))
	if err := cli.Connect(); err != nil {
		return "", errors.New(constants.HostInfoError)
	}

	//延迟注册关闭
	defer func(cli *utils.SSH) {
		_ = cli.Client.Close()
		_ = cli.Session.Close()
	}(cli)

	output, err := cli.Run(cmd)
	if err != nil {
		return "", err
	}
	return output, err
}

//批量主机执行命令

func ExecuteCMDList(HostIds []int, Cmd string) error {
	for _, id := range HostIds {
		fmt.Println("id>>> ", id)
		_, err := ExecuteCmd(id, Cmd)
		if err != nil {
			return err
		}
	}
	return nil
}
