package main

import (
	"fmt"
	"sync"
	"time"
)

func read() {
	defer wg.Done() //执行完毕后计数器减一
	fmt.Println("read start...")
	time.Sleep(time.Second * 3) //模拟I/O阻塞
	fmt.Println("read end...")
}

func listenMusci() {
	defer wg.Done()
	fmt.Println("listenMusci start...")
	time.Sleep(time.Second * 5)
	fmt.Println("listenMusci end...")
}

// 协程同步锁
var wg sync.WaitGroup

func main() {
	strart := time.Now().Unix()
	wg.Add(2) //一个协程加1,两个就是2
	go read() //并发执行
	go listenMusci()
	wg.Wait() //计数器为0继续执行
	end := time.Now().Unix()
	fmt.Println("执行完毕...", end-strart)
}
