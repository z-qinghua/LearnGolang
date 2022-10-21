// @program:     Eratosthenes2.0.go
// @file:        isprime.go
// @create:      2022-09-30 17:09
// @description:

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var prime_ls []int
	fmt.Print("输入正确的合数：")
	fmt.Scanln(&n)
	if n == 1 || n == 2 {
		fmt.Println(n, "是素数")
	} else {
		for i := 0; i <= n; i++ { //"="不能少
			prime_ls = append(prime_ls, i)
		}
		//fmt.Println(prime_ls)
		var k int = int(math.Sqrt(float64(n)))
		count := 0
		for i := 2; i < k+1; i++ {
			if prime_ls[i] != 0 {
				for j := i * i; j <= n; j += i {
					prime_ls[j] = 0
				}
			}
		}
		for i := 2; i < k+1; i++ {
			if prime_ls[i] != 0 && n%prime_ls[i] == 0 {
				count++
			}
		}
		if count != 0 {
			fmt.Println(n, "不是素数")
		} else {
			fmt.Println(n, "是素数")
		}
	}
}
