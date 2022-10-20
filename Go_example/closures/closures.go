// @program:     Go_example
// @file:        closures.go
// @create:      2022-10-11 11:32
// @description:

package main

import "fmt"

//这个`intSeq`函数返回另一个在`intSeq`函数体内定义的匿名函数
//这个返回的函数使用闭包的方式_隐藏_变量_`i`
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    //我们调用`intSeq`函数，将返回值（一个函数）赋给`nextInt`。
    //这个函数的值包含了自己的值`i`
    //这样在每次调用`nextInt`时都会更新`I`的值。
    nextInt := intSeq()

    //通过多次调用`nextInt`来看看闭包的效果
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    //为了确认这个状态对于这个特定的函数是唯一的，我们重新创建并测试一下
    newInts := intSeq()
    fmt.Println(newInts())
}
