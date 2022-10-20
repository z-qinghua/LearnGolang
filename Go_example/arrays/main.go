// @program:     Go_example
// @file:        main.go
// @create:      2022-10-09 17:20
// @description:

//在Go中，数组是一个具有固定长度且编号的元素序列。
package main

import "fmt"

func main() {
	//创建一个数组a来存放刚好5个int
	//元素的类型和长度都是数组类型的一部分
	//数组默认是零值，对于int数组来说也就是0
	var a [5]int
	fmt.Println("emp:", a)

	//我们可以使用array[index]=value语法来设置数组,指定位置的值，
	//或者用array[index]得到值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//使用内置函数len返回数组的长度
	fmt.Println("len:", len(a))

	//使用这个语法在一行内声明并初始化一个数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//数组是一维的，但是可以自由组合构造多维的数据结构
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
