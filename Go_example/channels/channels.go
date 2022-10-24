// @program:     Go_example
// @file:        channels.go
// @create:      2022-10-13 13:01
// @description:

//通道(channels)是连接多个Go协程的管道。
//你可以从一个Go协程将值发送到通道，然后在别的Go协程中接收。
package main

import "fmt"

func main() {
	//使用make(chan val-type)创建一个新的通道。
	//通道类型就是他们需要传递值的类型。
	message := make(chan string)

	//使用channel <- 语法 _发送(send)_一个新的值到通道中。
	//我们在一个新的Go协程中发送“ping"到上面创建的messages通道中
	go func() { message <- "ping" }()

	//使用<-channel语法从通道中_接收_一个值。
	//这里接收上面发送的ping消息并打印
	msg := <-message
	fmt.Println(msg)
}

//默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。默认通道是无缓冲的
