// @program:     Go_example
// @file:        multiple-return-value.go
// @create:      2022-10-11 10:25
// @description:

//GO内建多返回值支持。这个特性在Go语言中被经常用到，
//例如用来同时返回一个函数的结果和错误信息。
package main

import "fmt"

//(int, int)在这个函数中标志着这个函数返回2个int
func vals() (int, int) {
	return 3, 7
}

func main() {
	//这里我们通过_多赋值_操作来使用这两个不同的返回值。
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//如果你仅仅需要返回值的一部分的话，可以使用空白标识符`_`
	_, c := vals()
	fmt.Println(c)
}
