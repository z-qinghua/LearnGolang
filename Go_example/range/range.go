// @program:     Go_example
// @file:        range.go
// @create:      2022-10-10 16:17
// @description:

//range迭代各种各样的数据结构。
package main

import "fmt"

func main() {
    //使用range来对slice中的元素求和
    //对于数组也可以采用这种办法。
    nums := []int{1, 2, 3}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    //range在数组和slice中提供对每项的索引和值的访问
    //我们不需要索引时，使用_空白标识符'_'来忽略它，
    //有时候我们实际上是需要这个索引的。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    //range在map中迭代键值对
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    //range也可以只遍历map的键
    for k := range kvs {
        fmt.Println("key:", k)
    }
    //range在字符中迭代Unicode码点
    //第一个返回值是字符的起始字节位置，然后第二个是字符本身
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
