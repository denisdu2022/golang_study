package services

import (
	"bingotest01/application/model"
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
	if err = ctx.ShouldBindJSON(&host); err != nil {
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
