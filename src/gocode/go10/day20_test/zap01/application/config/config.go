package config

import (
	"encoding/json"
	"io/ioutil"
)

//全局日志配置结构体

type LogConfig struct {
	//日志级别
	Level string `json:"level"`
	//单个文件大小 单位MB
	MaxSize int `json:"max_size"`
	//文件保留天数
	MaxAge int `json:"max_age"`
	//文件保留个数
	MaxBackups int `json:"max_backups"`
	//是否启用压缩
	Compress bool `json:"compress"`
}

//全局配置结构体

type Config struct {
	//模式
	Mode string `json:"mode"`
	//服务启动端口
	Port int `json:"port"`
	//日志成员
	*LogConfig `json:"log"`
}

// 全局配置对象

var Conf = new(Config)

//初始化配置文件

func Init(filePath string) error {
	//读取json配置文件
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	//反序列读取的json配置文件并返回
	return json.Unmarshal(b, Conf)
}
