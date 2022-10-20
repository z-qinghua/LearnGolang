// @program:     Go_example
// @file:        structs.go
// @create:      2022-10-11 13:45
// @description:

//Go的结构体(struct)是带类型的字段(fields)集合。这在组织数据结构时很有用
package main

import "fmt"

//这里的person结构体包含了name和age两个字段
type person struct {
	name string
	age  int
}

func main() {
	//使用这个语法创建新的结构体元素
	fmt.Println(person{"bob", 20})

	//可以在初始化一个结构体元素时指定字段名字
	fmt.Println(person{name: "Alice", age: 30})

	//省略的字段将被初始化为0
	fmt.Println(person{name: "Fred"})

	//&前缀生成一个结构体指针
	fmt.Println(&person{name: "Ann", age: 40})

	//使用`.`来访问结构体字段
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	//也可以对结构体指针使用`.`，指针会被自动解引用
	sp := &s
	fmt.Println(sp.age)
	fmt.Println(sp)

	//结构体是可变的(mutable)
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(*sp)
}
