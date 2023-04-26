package database

import (
	"bingo_api/applcation/config"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"

	//在引入包的前面使用_ "github.xx/xx/xxx" 就是只使用包的初始化
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//定义全局Orm对象

var Orm *gorm.DB

//初始化数据库连接函数

func InitDB(cfgDB *config.DatabaseConfig) *gorm.DB {
	//定义错误对象
	var err error

	//定义数据库连接信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		cfgDB.Username,
		cfgDB.Password,
		cfgDB.Host,
		cfgDB.Port,
		cfgDB.Database,
		cfgDB.Charset)

	//连接数据库
	Orm, err = gorm.Open(cfgDB.Driver, dsn)
	//判断err,将连接数据库失败的错误信息写到日志记录
	if err != nil {
		//调用zap.S()函数,日志级别为Error 将错误信息记录到日志中
		zap.S().Errorf("数据库连接失败,失败原因是: %s", err.Error())
		//把错误抛出
		panic(err.Error())
	}

	//设置数据库最大连接数
	Orm.DB().SetMaxOpenConns(cfgDB.MaximumConn)
	//设置数据库最大空闲连接数
	Orm.DB().SetMaxIdleConns(cfgDB.MaximumFreeConn)
	//设置数据库连接的最大生命周期
	Orm.DB().SetConnMaxLifetime(time.Duration(cfgDB.TimeOut))

	//判断Orm错误
	if Orm.Error != nil {
		//记录日志
		zap.S().Errorf("数据库错误,错误信息是: %v", Orm.Error)
		panic(Orm.Error)
	}

	return Orm
}
