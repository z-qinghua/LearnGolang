// @program:     Go_example
// @file:        slices.go
// @create:      2022-10-09 22:32
// @description:

//slice是Go中一个关键的数据类型，是一个比数组更加强大的序列接口
package main

import (
	"fmt"
)

func main() {
	//与数组不同，slice的类型仅由它所包含的元素决定（不需要元素的个数）
	//要创建一个长度非零的空slice，需要使用内建的方法make
	//这里我们创建了一个长度为3的string类型slice（初始化为零值）
	s := make([]string, 3)
	fmt.Println("emp:", s) //初始值为空字符

	//可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	//返回slice的长度
	fmt.Println("len:", len(s))

	//除了基本操作外，slice支持比数组更丰富的操作
	//其中一个是内建的append，它返回一个包含了一个或者多个新值的slice
	//注意由于append可能返回新的slice，我们需要接受其返回值
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//slice也可以被copy，这里我们创建一个空的和s有相同长度的slice c,
	//并将s复制给c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	//slice支持通过slice[low:high]语法进行切片操作
	l := s[2:5]
	fmt.Println("sl1:", l)

	//这个slice从s[0]切片到s[5](不包含)
	l = s[:5]
	fmt.Println("sl2:", l)

	//这个slice从s[2]（包含）开始切片
	l = s[2:]
	fmt.Println("sl3:", l)

	//我们可以在一行代码中声明并初始化一个slice变量
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//slice可以组成多维数据结构
	//内部的slice长度可以不一致，这和多维数组不同
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

//slice和数组是不同的类型，但是通过fmt.Println打印效果类似
