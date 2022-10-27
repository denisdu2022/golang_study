package main

import (
	"fmt"
	"net"
)

func main() {

	//1.确立网络三要素,建立服务
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	//错误信息处理
	if err != nil {
		fmt.Println("错误信息是: ", err)
		return
	}
	fmt.Println("listen", listen)

	//延迟注册关闭
	defer listen.Close()

	for true {
		//2.等待用户连接,没有连接的客户端保持阻塞
		fmt.Println("server is waiting...")
		conn, err := listen.Accept()
		//错误信息处理
		if err != nil {
			fmt.Println("错误信息是: ", err)
			return
		}
		fmt.Println("conn", conn)

		//接收客户端发来的请求
		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		fmt.Println("n", n)

		//响应客户端请求
		res := `HTTP/1.1 200 ok\r\n\r\n<h1>hello Web!!!</h1> <img src="https://gd-hbimg.huaban.com/b455e506e3162c8bbdb4545b69c4dab010592ecd8f99b-0bT5lq_fw658">`
		conn.Write([]byte(res))
	}

}
