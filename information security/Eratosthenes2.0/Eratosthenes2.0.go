// @program:     information security
// @file:        Eratosthenes2.0.go
// @create:      2022-09-29 21:16
// @description:

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var prime_ls []int
one:
	fmt.Print("输入正确的合数：")
	fmt.Scanln(&n)
	if n == 0 || n == 1 {
		goto one
	}
	for i := 0; i <= n; i++ { //"="不能少
		prime_ls = append(prime_ls, i)
	}
	//fmt.Println(prime_ls)
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if prime_ls[i] != 0 {
			for j := i * i; j <= n; j += i {
				prime_ls[j] = 0
			}
		}
	}
	var count int
	var prime []int
	for _, value := range prime_ls {
		if value != 0 && value != 1 {
			count++
			prime = append(prime, value)
		}
	}
	fmt.Println(count, prime)
}
