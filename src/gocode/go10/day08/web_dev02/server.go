package main

import (
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>hello Web!!!</h1> <img src="https://gd-hbimg.huaban.com/53aa44421080134f9465b7b8fe06f01734464272200748-OCJsnX_fw658">`)

}

func main() {
	//2.路由函数返回
	http.HandleFunc("/", foo)

	//1.开启服务端监听
	http.ListenAndServe(":8000", nil)

}
