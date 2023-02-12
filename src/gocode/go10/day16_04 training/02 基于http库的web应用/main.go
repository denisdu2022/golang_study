package main

import (
	"fmt"
	"net/http"
)

//首页路由函数

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>这是首页.....</h1?")
}

// 图片页面路由函数
func img(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>樱花图片</h1><img src=\"https://cdn.pixabay.com/photo/2023/01/14/19/50/flower-7718952_1280.jpg\">")
}

func main() {

	fmt.Println("Server is waiting!!!")

	//2.首页路由
	http.HandleFunc("/", index)

	//3.图片页面
	http.HandleFunc("/img", img)

	//1.开启服务端监听
	http.ListenAndServe(":8080", nil)

}
