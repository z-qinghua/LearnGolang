// @program:     Eratosthenes2.0.go
// @file:        euclid.go
// @create:      2022-10-01 12:38
// @description:

package main

import "fmt"

func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func main() {
	var a, flag int
	fmt.Print("输入一个正整数：")
	fmt.Scan(&a)
	if a == 1 || a == 2 {
		fmt.Println(a, "不能被非零整数整除")
	} else {
		for i := 2; i < a; i++ {
			if gcd(a, i) == 1 {
				flag++
				continue
			} else {
				fmt.Println(a, "能被非零整数整除")
				break
			}
		}
		if flag == a-2 {
			fmt.Println(a, "不能被非零整数整除")
		}
	}
}
