package main

import (
	"fmt"
	"net/http"
)

//首页函数

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>这是首页!!!</h1>`)
}

//图片页面函数

func img(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<img src="https://images.pexels.com/photos/2662116/pexels-photo-2662116.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2">`)
}

func main() {

	//首页函数
	http.HandleFunc("/", index)

	//图片页面函数
	http.HandleFunc("/img", img)

	//1.开启http监听
	http.ListenAndServe(":8000", nil)
}
