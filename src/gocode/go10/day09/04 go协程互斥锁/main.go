package main

import (
	"fmt"
	"sync"
	"time"
)

func read() {
	//defer wg.Done() //执行完毕后计数器减一
	fmt.Println("read start...")
	time.Sleep(time.Second * 3) //模拟I/O阻塞
	fmt.Println("read end...")
}

func listenMusci() {
	//defer wg.Done()
	fmt.Println("listenMusci start...")
	time.Sleep(time.Second * 5)
	fmt.Println("listenMusci end...")
}

// 协程同步锁
var wg sync.WaitGroup

// 互斥锁
var lock sync.Mutex

var x = 0

func add() {
	defer wg.Done()
	lock.Lock() //开启锁
	x++
	println(x)
	lock.Unlock() //关闭锁
	time.Sleep(time.Millisecond * 100)

}

func main() {
	//add()
	//fmt.Println(x) //1

	//for i:=0; i < 100; i++{
	//	add()
	//}
	//fmt.Println(x) //100

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
