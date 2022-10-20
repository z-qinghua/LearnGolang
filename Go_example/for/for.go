// @program:     Go_example
// @file:        for.go
// @create:      2022-10-09 14:49
// @description:

//for循环
//for是Go中唯一的循环结构。这里有for循环的三个基本使用方式。
package main

import (
	"fmt"
)

func main() {
	//最基础的方式，单个循环条件。
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//经典的初始/条件/后续`for`循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//不带条件的`for`循环将一直重复执行，直到在循环体内
	//使用了break或者return来跳出循环
	for {
		fmt.Println("loop")
		break
	}

	//你也可以使用continue来跳到下一个循环迭代
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue //偶数跳出，不输出
		}
		fmt.Println(n)
	}
}
