package model

import (
	"bingotest01/application/database"
	"github.com/jinzhu/gorm"
)

//主机类别实体

type HostCategory struct {
	gorm.Model
	//唯一索引 unique_index
	Name string `json:"name" gorm:"type:varchar(255);unique_index"`
}

//查询

type HostCategoryInstance struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

//设置表名

func (category HostCategory) TableName() string {
	return "host_category"
}

//添加主机类别

func (category *HostCategory) Insert() error {
	err := database.Orm.Create(&category).Error
	return err
}

//获取主机分类列表

func (category HostCategory) GetAll() ([]HostCategoryInstance, error) {
	//获取主机分类列表
	var categories []HostCategoryInstance
	err := database.Orm.Table(category.TableName()).Order("id DESC").Select("id,name").Scan(&categories).Error
	return categories, err
}

//根据指定ID获取主机类别

func (category HostCategory) GetOneById(id uint) error {
	err := database.Orm.First(&category, id).Error
	return err
}

//主机信息模型

type Host struct {
	gorm.Model
	//主机名
	Name string `json:"name" gorm:"type:varchar(255)"`
	//IP地址
	IpAddr string `json:"ip_addr" gorm:"type:varchar(255)"`
	//端口
	Port uint `json:"port" gorm:"type:int"`
	//用户名
	Username string `json:"username" gorm:"varchar(255)"`
	//密码
	Password string `json:"password,omitempty" gorm:"varchar(255)"`
	//备注
	Remark string `json:"remark,omitempty" gorm:"varchar(255)"`
	//关联ID
	HostCategoryID uint `json:"host_category_id"`
	//关联关系
	HostCategory HostCategory `json:"host_category" gorm:"foreignKey:HostCategoryID;references:ID"`
}

type HostInstance struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	IpAddr       string `json:"ip_addr"`
	Port         uint   `json:"port"`
	Username     string `json:"username"`
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	Remark       string `json:"remark"`
}

//设置表名

func (host Host) TableName() string {
	return "host_info"
}

//添加主机

func (host *Host) Insert() error {
	err := database.Orm.Create(&host).Error
	return err
}

//获取所有主机的列表

func (host Host) GetAll(name string, hostCategoryId uint, IpAddr string) ([]Host, error) {
	//实例化hosts对象
	var hosts []Host
	query := database.Orm.Table(host.TableName())
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if hostCategoryId > 0 {
		query = query.Where("host_category_id = ?", hostCategoryId)
	}
	if IpAddr != "" {
		query = query.Where("ip_addr LIKE ?", "%s"+IpAddr+"%s")
	}

	err := query.Order("id DESC").Preload("HostCategory").Find(&hosts).Error
	return hosts, err
}

//根据指定ID获取主机

func (host *Host) GetOneById(id uint) error {
	err := database.Orm.First(&host, id).Error
	return err
}

//删除主机

func (host *Host) Delete() (err error) {
	err = database.Orm.Delete(&host).Error
	return err
}
