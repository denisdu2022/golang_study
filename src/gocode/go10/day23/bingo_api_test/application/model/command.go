package model

import (
	"bingotest01/application/database"
	"github.com/jinzhu/gorm"
)

//指令模板类别

type CommandCategory struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255);unique_index"`
}

//指令模板类别实例

type CommandCategoryInstance struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

//设置表名

func (category CommandCategory) TableName() string {
	return "command_category"
}

//添加指令模板类别

func (category *CommandCategory) Insert() error {
	err := database.Orm.Create(&category).Error
	return err
}

//获取指令模板分类列表

func (category CommandCategory) GetAll() ([]CommandCategoryInstance, error) {
	//实例化categories对象
	var categories []CommandCategoryInstance
	err := database.Orm.Table(category.TableName()).Order("id DESC").Select("id,name").Scan(&categories).Error
	return categories, err
}

//根据指定ID获取指令模板类别

func (category *CommandCategory) GetOneById(id uint) error {
	err := database.Orm.First(&category, id).Error
	return err
}

//指令模板

type Command struct {
	gorm.Model
	Name              string          `json:"name" gorm:"type:varchar(255)"`
	Content           string          `json:"content" gorm:"type:text"`
	CommandCategoryId uint            `json:"command_category_id"`
	CommandCategory   CommandCategory `json:"command_category" gorm:"foreignKey:CommandCategoryID;references:ID"`
}

//指令模板实例

type CommandInstance struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
}

//设置表名

func (command Command) TableName() string {
	return "command_template"
}

//添加指令模板

func (command *Command) Insert() error {
	err := database.Orm.Create(&command).Error
	return err
}

//获取所有指令模板列表

func (command Command) GetAll(name string, commandCategoryId uint) ([]Command, error) {
	//实例化command对象
	var commands []Command
	//数据库查询
	query := database.Orm.Table(command.TableName())

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if commandCategoryId > 0 {
		query = query.Where("command_category_id = ?", commandCategoryId)
	}

	err := query.Order("id DESC").Preload("CommandCategory").Find(&commands).Error
	return commands, err
}

//根据指定ID获取指令模板

func (command *Command) GetOneById(id uint) error {
	err := database.Orm.Preload("CommandCategory").First("&command", id).Error
	return err
}

//删除指令模板

func (command *Command) Delete() (err error) {
	err = database.Orm.Delete(&command).Error
	return err
}
