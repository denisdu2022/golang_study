package api //包名与目录名同名
import (
	"fmt"
	"mysite/db"
)

func RestfulAPI() {
	db.HandleMySql()
	fmt.Println("RestfulAPI:MySql数据接口...")
}
