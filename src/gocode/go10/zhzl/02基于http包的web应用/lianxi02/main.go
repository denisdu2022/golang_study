package main

import (
	"fmt"
	"net/http"
)

// 首页函数
func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello index web!</h1>")
}

// 图片函数
func getImg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>樱花图片</h1><img src=\"https://cdn.pixabay.com/photo/2023/01/14/19/50/flower-7718952_1280.jpg\">")
}

func main() {
	//基于http包的web应用

	fmt.Println("Server is Waiting!!!")

	//首页路由
	http.HandleFunc("/", getIndex)

	//图片路由
	http.HandleFunc("/img", getImg)

	//1.开启服务端监听
	http.ListenAndServe("127.0.0.1:8080", nil)
}
