// @program:     Eratosthenes2.0.go
// @file:        relatively prime.go
// @create:      2022-09-30 20:46
// @description:

package main

import "fmt"

//递归
func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func main() {
	var m, a int
	fmt.Print("输入两个正整数：")
	fmt.Scan(&m, &a)
	if m < a {
		c := a
		a = m
		m = c
	}
	if gcd(m, a) == 1 {
		fmt.Println("a与m互素")
	} else {
		fmt.Println("a与m不互素")
	}
}
