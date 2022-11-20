package main

import (
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<h1>hello web</h1>")
	if err != nil {
		fmt.Println("错误信息是: ", err)
		return
	}
}

func img(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, `<img src="https://gd-hbimg.huaban.com/08c37af52b5735e3a1ae950708495cefe2b6e442b6d7b-wUj1xE_fw658">`)
	if err != nil {
		fmt.Println("错误信息是: ", err)
		return
	}
}

func main() {

	//02 路由
	http.HandleFunc("/", foo)
	http.HandleFunc("/img", img)
	//01 开启http协议监听
	err := http.ListenAndServe(":8000", nil)
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
		return
	}
}
