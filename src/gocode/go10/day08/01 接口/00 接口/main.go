package main

import "fmt"

// 定义接口类
// 定义一个支付的接口类
type Pay interface {
	//方法名  返回值
	pay() string
}

// Alipay
type Alipay struct {
}

// Alipay的方法接收器
func (a Alipay) pay() string {
	fmt.Println("Alipay支付...")
	return "阿里支付"
}

// Weixinpay
type Weixinpay struct {
}

// Weixinpay的方法接收器
func (w Weixinpay) pay() string {
	fmt.Println("Weixinpay支付...")
	return "微信支付"
}

// 订单类
func oder(p Pay) {
	p.pay()
}

func main() {

	//实例化Alipay对象
	var p Pay
	p = Alipay{}
	p.pay() //Alipay支付...

	//Weixinpay对象
	p = Weixinpay{}
	p.pay() //Weixinpay支付...

	//接收返回值
	ret := p.pay()
	fmt.Println(ret) //微信支付

	fmt.Println("--------------------------------------------")
	oder(Alipay{})
	oder(Weixinpay{})
	/*
		输出结果
		Alipay支付...
		Weixinpay支付...
	*/

}
