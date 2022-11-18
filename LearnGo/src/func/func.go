// @program:     LearnGo
// @file:        func.go
// @create:      2022-09-19 11:59
// @description:

package main

import "fmt"

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation:" + op)
	}
}

func div(a, b int) (q, r int) { //函数可以返回多个值
	//q = a / b
	//r = a % b

	return a / b, a % b
}

//交换数值
/*
func swap(a, b *int) {
	*a, *b = *b, *a
}
*/
func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(eval(3, 4, "-"))
	q, r := div(3, 4)
	fmt.Println(q, r)

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
