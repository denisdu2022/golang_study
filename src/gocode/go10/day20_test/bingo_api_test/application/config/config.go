package config

import (
	"encoding/json"
	"io/ioutil"
)

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
