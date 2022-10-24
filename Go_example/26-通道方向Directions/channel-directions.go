// @program:     Go_example
// @file:        channel-directions.go
// @create:      2022-10-24 14:20
// @description:

//当使用通道作为函数的参数时，可以指定这个通道是不是只用来发送或者接收值。
//这个特性提升了程序的类型安全性。
package main

import "fmt"

//`ping`函数定义了一个只允许发送数据的通道。
//尝试使用这个通道来接收数据将会得到一个编译错误
func ping(ping chan<- string, msg string) {
	ping <- msg
}

//`pong`函数定义了一个只允许接收来自通道(ping)的数据。
func pong(ping <-chan string, pongs chan<- string) {
	msg := <-ping
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
