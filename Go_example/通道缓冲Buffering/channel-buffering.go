// @program:     Go_example
// @file:        channel-buffering.go
// @create:      2022-10-24 09:24
// @description:

//通道缓冲Buffering
//默认情况下，通道是无缓冲的，这意味这只有对应的接收（-<chan）通道准备接收时，
//才允许发送（chan<-）。可缓存通道允许在没有对应接收方的情况下，缓存限定量的值。

package main

import "fmt"

func main() {
	//创建一个字符串通道，最多允许缓存2个值。
	messages := make(chan string, 2)

	//由于此通道是缓冲的，因为我们可以将这些值发送到通道中，而不需要相应的并发接收
	messages <- "buffered"
	messages <- "channel"

	//然后我们可以像前面一样接收这两个值
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
