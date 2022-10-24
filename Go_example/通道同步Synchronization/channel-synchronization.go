// @program:     Go_example
// @file:        channel-synchronization.go
// @create:      2022-10-24 11:07
// @description:

//通道同步：
//可以使用通道来同步Go协程间的执行状态。这里是使用阻塞的接收方式来
//等待一个Go协程的运行结束

package main

import (
	"fmt"
	"time"
)

//这是一个我们将要在Go协程中运行的函数。done通道将被用于通知其它Go协程
//这个函数已经工作完毕。
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//发送一个值来通知我们已经完工了
	done <- true
}

func main() {
	fmt.Println("main() started")

	//运行一个worker Go协程，并给予用于通知的通道
	done := make(chan bool, 1)
	go worker(done) //worker函数以go关键词以协程方式运行它
	//此时两个协程并且正在调度运行的是main goroutine

	//程序将在接收通道中worker发出的通知前一直阻塞。
	<-done                 //从done读取一些数据
	fmt.Println("channel") //等worker运行结束才会执行
	//如果把<-done这行代码移除，程序会在worker还没开始运行就结束
}
