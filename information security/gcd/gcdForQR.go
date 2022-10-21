// @program:     Eratosthenes.go
// @file:        gcdForQR.go
// @create:      2022-10-17 20:45
// @description:

package main

import "fmt"

func gcd() {
	a, b := 0, 0
	//输入两个数
	fmt.Printf("输入两个数：")
	fmt.Scanln(&a, &b)
	var q, r int
	for {
		if a < b {
			r = a
			break
		} else {
			a = a - b
			q = q + 1
		}
	}
	fmt.Printf("q=%d  r=%d\n", q, r)
}

func main() {
	gcd()
}
