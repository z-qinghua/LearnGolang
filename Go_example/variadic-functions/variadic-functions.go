// @program:     Go_example
// @file:        variadic-functions.go
// @create:      2022-10-11 10:53
// @description:

//可变参数函数。在调用时可以任意数量的参数。
//例如，fmt.Println是一个常见的变参函数。
package main

import "fmt"

//这个函数接受任意数目的`int`作为参数
func sum(nums ...int) {
	fmt.Print(nums, " ") //输出全部参数
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	//变参函数使用常规的调用方式，传入独立的参数。
	sum(1, 2)
	sum(1, 2, 3)

	//如果有一个含有多个值的slice，想把它们作为参数使用
	//你要这样调用`func(slice...)`
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
