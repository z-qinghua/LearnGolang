// @program:     information security
// @file:        Eratosthenes2.0.go
// @create:      2022-09-29 15:49
// @description:

package main

import "fmt"

func select_primes(prime []int, factor int) {
	for i, value := range prime {
		if value != 0 && value != factor {
			if prime[i]%factor == 0 {
				prime[i] = 0
			}
		}
	}
}

func main() {
	var n int
	fmt.Print("输入合数：")
	fmt.Scanln(&n)
	var prime_ls []int
	for i := 2; i <= n; i++ {
		prime_ls = append(prime_ls, i)
	}
	for i := 0; i < len(prime_ls); i++ {
		if prime_ls[i] != 0 {
			select_primes(prime_ls, prime_ls[i])

		}
	}
	var count int
	var arr []int
	for _, value := range prime_ls {
		if value != 0 {
			count++
			arr = append(arr, value)
		}
	}
	fmt.Println("素数个数为:", count)
	fmt.Println(arr)
}
