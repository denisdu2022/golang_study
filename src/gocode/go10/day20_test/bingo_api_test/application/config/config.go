package config

import (
	"encoding/json"
	"io/ioutil"
)

//DatabaseConfig数据库配置

type DatabaseConfig struct {
	//数据库驱动
	Driver string `json:"driver"`
	//数据库主机
	Host string `json:"host"`
	//数据库端口
	Port string `json:"port"`
	//库
	Database string `json:"database"`
	//用户名
	Username string `json:"username"`
	//密码
	Password string `json:"password"`
	//字符集
	Charset string `json:"charset"`
	//最大连接数
	MaximumConn int `json:"maximum_connection"`
	//空闲连接数
	MaximumFreeConn int `json:"maximum_free_connection"`
	//连接超时
	TimeOut int `json:"timeout"`
}

//日志对象结构体

type LogConfig struct {
	Level      string `json:"level"`       //日志级别
	Filename   string `json:"filename"`    //日志路径
	MaxSize    int    `json:"maxsize"`     //单个日志文件大小
	MaxAge     int    `json:"max_age"`     //日志周期
	MaxBackups int    `json:"max_backups"` //日志备份的数量
}

//Config 整个项目的配置

type Config struct {
	Mode string `json:"mode"`
	Port int    `json:"port"`
	//新增日志成员
	*LogConfig `json:"log"`
	//数据库配置项
	*DatabaseConfig `json:"database"`
}

//配置文件解析
//先读取config.json文件,在反序列化config.json

//Conf全局配置对象

var Conf = new(Config)

//初始化配置文件:从指定文件加载配置文件

func Init(filePath string) error {
	//filePath 配置文件json
	//ioutil读取配置文件

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	//如果正常读取到配置文件,则返回反序列化后的文件信息
	return json.Unmarshal(b, Conf)
}
