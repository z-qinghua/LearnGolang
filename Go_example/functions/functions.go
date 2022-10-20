// @program:     Go_example
// @file:        functions.go
// @create:      2022-10-11 10:08
// @description:

package main

import "fmt"

//这里是一个函数，接受两个int并且以int返回他们的和
func plus(a int, b int) int {
    //Go语言需要明确的返回，不会自动返回最后一个表达式的值
    return a + b
}

//当多个连续的参数为同样类型时，最多可以仅声明最后一个参数类型
//而忽略之前相同类型参数的类型声明
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {
    //通过name(args)来调用函数
    res := plus(1, 2)
    fmt.Println("1+2 = ", res)

    res1 := plusPlus(1, 2, 3)
    fmt.Println("1+2+3 = ", res1)
}
