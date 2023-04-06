package config

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("打印传进来的json数据>>>", string(b))
	return json.Unmarshal(b, Conf)
}
