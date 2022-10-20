// @program:     Go_example
// @file:        if_else.go
// @create:      2022-10-09 15:17
// @description:

//if_else:在Go语言中没有三目运算符，即使最基本的条件判断仍需要完整的if语句
package main

import (
	"fmt"
)

func main() {
	//这里是一个基本的例子
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//可以不要else只用if语句
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//在条件语句之前可以有一个声明语句；
	//这里声明的变量可以在所有的条件分支中山使用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
