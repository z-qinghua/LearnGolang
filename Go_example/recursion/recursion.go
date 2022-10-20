// @program:     Go_example
// @file:        recursion.go
// @create:      2022-10-11 11:43
// @description:

//递归：一下是经典的阶乘
package main

import "fmt"

//fact函数在到达fact(0)之前一直调用自身
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
