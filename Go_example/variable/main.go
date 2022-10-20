// @program:     Go_example
// @file:        main.go
// @create:      2022-10-08 16:03
// @description: 

//变量：被显示声明，并可以被编译器用来检查函数调用时的类型正确性。
package main

import "fmt"

func main() {
    //`var` 声明1一个或者多个变量。
    var a = "initial"
    fmt.Println(a)

    //可以一次性声明多个变量。
    var b, c int = 1, 2
    fmt.Println(b, c)

    //Go将自动推断已经初始化的变量类型。
    var d = true
    fmt.Println(d)

    //声明后却没有给相应初始值时，变量会初始化为“零值”
    var e int
    fmt.Println(e)

    //`:=`语法是声明并初始化变量的简写。
    f := "short"
    fmt.Println(f)
}
