package main

import "github.com/gin-gonic/gin"

// 结构体
type Student struct {
	Name   string
	Age    int
	Gender string
	Email  string
	Addr   string
}

func index(ctx *gin.Context) {
	//构建Go语言下的各类数据
	//变量
	var name = ""
	var age = 22
	var gender = "male"
	var email = "qwqwq@qq.com"
	//切片
	var books18 = []string{"三国演义", "红楼梦", "西游记", "水浒传", "资治通鉴"}
	var books1 = []string{"三遂平妖", "葫芦娃", "喜羊羊与灰太狼之兔年顶呱呱"}
	var books = []string{"金瓶梅", "聊斋", "剪灯新话"}
	//取第一本书的实现方法,不利于维护
	//var books01 = books18[0]

	//map对象
	var stuMap = map[string]interface{}{
		"name":   "hb",
		"age":    22,
		"gender": "male",
	}

	//结构体对象
	stuStruct := Student{Name: "zhangsna", Age: 18}

	var students = []Student{{"李四", 21, "male", "111@qq.com", "河北省张家口市"}, {"王五", 24, "male", "222@qq.com", "河北省承德市"}, {"赵六", 25, "male", "333@qq.com", "河北省石家庄"}}

	ctx.HTML(200, "index.html", gin.H{
		"name":    name,
		"age":     age,
		"gender":  gender,
		"email":   email,
		"books18": books18,
		"books1":  books1,
		//"books01": books01,
		"books":     books,
		"stuMap":    stuMap,
		"stuStruct": stuStruct,
		"students":  students,
	})
}

func main() {
	//获取路由对象
	r := gin.Default()
	//加载模版文件
	r.LoadHTMLGlob("templates/*")
	//设置静态文件资源窗口
	r.Static("/static", "./static")

	r.GET("/", index)

	//启动
	r.Run()

}
