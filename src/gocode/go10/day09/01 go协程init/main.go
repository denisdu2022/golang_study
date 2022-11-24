package main

import (
	"fmt"
	"time"
)

func read() {
	fmt.Println("read start...")
	time.Sleep(time.Second * 3) //模拟I/O阻塞
	fmt.Println("read end...")
}

func listenMusci() {
	fmt.Println("listenMusci start...")
	time.Sleep(time.Second * 5)
	fmt.Println("listenMusci end...")
}

func main() {
	//start := time.Now().Unix()
	go read() //并发执行
	go listenMusci()
	//一个主线程,两个子线程/协程
	//end := time.Now().Unix()
	time.Sleep(time.Second * 10)
	//fmt.Println(end - start)  //0 主死随从

}
