// @program:     Go_example
// @file:        timers.go
// @create:      2022-10-13 09:11
// @description:

//定时器：
//我们常常需要在后面一个时刻运行Go代码，或者某段时间间隔内重复运行。
//GO的内置定时器和打点器特性让这些很容易实现。
package main

import (
    "fmt"
    "time"
)

func main() {

    //定时器表示在未来某一时刻的独立事件。
    //你告诉定时器需要等待的时间，然后他将提供一个用于通知的通道。
    //这里的定时器将等待 2 秒
    timer1 := time.NewTimer(time.Second * 2)

    //<-time.C 直到这个定时器的通道 C 明确的发送了
    //定时器失效的值之前，将一直阻塞
    <-timer1.C
    fmt.Println("Timer 1 expired")

    //如果你需要的仅仅是单纯的等待，你需要使用time.Sleep
    //定时器有用的原因之一就是你可以在定时器失效之前，取消这个定时器。
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()    //bool类型
    fmt.Printf("%T\n", stop2) //输出类型

    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
