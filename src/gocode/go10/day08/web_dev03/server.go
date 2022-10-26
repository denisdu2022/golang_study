package main

import (
	"fmt"
	"net/http"
)

// 页面函数
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>hello Web!!!</h1>`)

}

func img(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h2>tupian</h2> <img src="https://gd-hbimg.huaban.com/08c37af52b5735e3a1ae950708495cefe2b6e442b6d7b-wUj1xE_fw658">`)
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/img", img)
	//1.开启http监听
	http.ListenAndServe(":8000", nil)

}
