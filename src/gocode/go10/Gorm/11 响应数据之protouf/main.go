package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

// protobuf函数
func protobuf(ctx *gin.Context) {
	//定义reps切片
	reps := []int64{int64(9), int64(55)}

	//自定义数据
	label := "今天天气真不错!!!"
	//传protobuf格式数据
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	fmt.Println(data)
	ctx.ProtoBuf(200, data)
}

func main() {
	//获取路由对象
	r := gin.Default()

	//响应protobuf格式数据
	//路由
	r.GET("/protobuf", protobuf)

	//启动
	r.Run()
}
