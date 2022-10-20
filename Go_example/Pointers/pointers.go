// @program:     Go_example
// @file:        pointers.go
// @create:      2022-10-11 13:09
// @description:

//Go支持指针，允许在程序中通过引用传递值或者数据结构
package main

import "fmt"

//我们将通过两个函数：zeroval和zeroptr来比较指针和值类型的不同
//zeroval有一个int型参数，所以使用值传递
//zeroval将从调用它的那个函数中得到一个ival形参的拷贝
func zeroval(ival int) {
	ival = 0
}

//zeroptr有一和上面不同的`*int`参数，意味着它用了一个int指针
//函数体内的`*iptr`接着_引用_这个指针，从它内存地址得到这个地址对应的当前值
//对一个解引用的指针赋值将会改变这个指针的真实地址的值
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	//通过&i语法来取得i的内存地址，即指向i的指针
	zeroptr(&i) //变量内存地址的引用
	fmt.Println("zeroptr:", i)

	//指针也可以被打印
	fmt.Println("pointer:", &i)
}
