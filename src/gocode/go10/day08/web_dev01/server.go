package main

import (
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>hello Web!!!</h1>`)

}

func main() {

	//1.开启服务端监听
	http.ListenAndServe(":8000", nil)

	//2.路由函数返回
	http.HandleFunc("/hi", foo)

}
