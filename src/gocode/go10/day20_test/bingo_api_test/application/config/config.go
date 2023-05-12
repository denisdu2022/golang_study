package config

import (
	"encoding/json"
	"io/ioutil"
)

//Config 整个项目的配置

type Config struct {
	Mode string `json:"mode"`
	Port int    `json:"port"`
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
