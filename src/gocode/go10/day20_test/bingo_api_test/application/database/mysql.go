package database

import (
	"bingotest01/application/config"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
	//在引入包的前面使用_ "github.xx/xx/xxx" 就是只使用包的初始化
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//全局Orm对象

var Orm *gorm.DB

//数据库初始化函数

func InitDB(cfg *config.DatabaseConfig) *gorm.DB {
	//获取链接信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
	)

	var err error
	//获取Orm连接对象
	Orm, err = gorm.Open(cfg.Driver, dsn)
	if err != nil {
		zap.S().Errorf("数据库连接失败,错误信息是: %v", err.Error())
		panic(err.Error())
	}

	//设置最大连接数
	Orm.DB().SetMaxOpenConns(cfg.MaximumConn)
	//设置最大空闲连接数
	Orm.DB().SetMaxIdleConns(cfg.MaximumFreeConn)
	//设置每个数据库连接的最大生命周期
	Orm.DB().SetConnMaxLifetime(time.Duration(cfg.TimeOut))
	if Orm.Error != nil {
		zap.S().Errorf("数据库错误吗,错误信息是: %v", Orm.Error)
		panic(Orm.Error)
	}

	return Orm
}
