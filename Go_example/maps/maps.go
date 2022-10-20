// @program:     Go_example
// @file:        maps.go
// @create:      2022-10-10 09:38
// @description:

//map是Go内置关联数据类型（在一些语言中称为哈希hash或者字典dict）
package main

import "fmt"

func main() {
	//要创建一个空map，需要使用内建的make：make(map[key type]value type)
	m := make(map[string]int)
	fmt.Println(m)

	//使用典型的make[key] = value语法设置键值对
	m["k1"] = 7
	m["k2"] = 13

	//使用例如fmt.Println来打印一个map会输出所有键值对
	fmt.Println(m)

	//使用name[key]来获取一个键的值
	v1 := m["k1"]
	fmt.Println("k1:", v1)

	//当对一个map调用内建的len时，返回的是键值对数目
	fmt.Println("len of map is ", len(m))

	//内建的delete可以从一个map中移除键值对
	delete(m, "k2")
	fmt.Println("map:", m)
	fmt.Println("len of map is ", len(m))

	//当从一个map中取值时，可选第二返回值指示这个键
	//是否在这个map中，这可以用来消除键不存在和键有零值
	//不需要值时，用_空白标识符忽略
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//可以通过以下语法在同一行声明和初始化一个新的map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map2:", n)
}

//注意一个 map 在使用 fmt.Println 打印的时候，是以 map[k:v k:v] 的格式输出的。
