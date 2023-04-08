package config

import (
	"encoding/json"
	"io/ioutil"
)

//外界go开发项目: 配置文件都是json , yaml , ini

//配置文件解析
//先读取config.json文件
//在反序列化config.json

//Config整个项目的配置

type Config struct {
	Mode string `json:"mode"`
	Port int    `json:"port"`
	//新增日志成员
	*LogConfig `json:"log"`
}

type LogConfig struct {
	Level      string `json:"level"`       //日志级别
	Filename   string `json:"filename"`    //日志路径
	MaxSize    int    `json:"maxsize"`     //单个日志文件的大小
	MaxAge     int    `json:"max_age"`     //日志周期
	MaxBackups int    `json:"max_backups"` //日志的备份数量
}

//Conf全局配置变量

var Conf = new(Config) //声明 var Conf Config

//Init初始化配置;从指定文件加载配置文件

func Init(filePath string) error {
	//filePath 配置文件json文件的路径

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	//fmt.Println("打印传进来的json数据>>>", string(b))
	return json.Unmarshal(b, Conf)
}
