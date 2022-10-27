package main

import (
	"fmt"
	"sync"
	"time"
)

// 生产者
func producer(ch chan<- int) { //chan<- 单向管道
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- i //把i的值写入管道
		fmt.Println("生产了消息:", i)
		time.Sleep(time.Second)

	}
	//生成完了关闭通道,解决了没有写数据的协程就会报错
	close(ch)

}

// 消费者
func consumer(ch <-chan int) {
	defer wg.Done()

	//for i := 0; i < 10; i++ {
	//	i := <-ch
	//	fmt.Println("消费了消息", i)
	//}

	////没有写数据的协程就会报错
	//for true {
	//	i := <-ch
	//	fmt.Println("消费了消息", i) //fatal error: all goroutines are asleep - deadlock!
	//}

	//遍历取值
	for i := range ch {
		fmt.Println("消费了消息", i)
	}

}

// 声明同步锁
var wg sync.WaitGroup

func main() {
	//生产者消费者模型
	ch := make(chan int, 100)

	wg.Add(2) //2个协程

	go producer(ch)
	go consumer(ch)

	wg.Wait()

	fmt.Println("process end ....")

}
