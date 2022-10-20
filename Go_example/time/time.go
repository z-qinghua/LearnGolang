// @program:     Go_example
// @file:        time.go
// @create:      2022-10-13 22:36
// @description:

//时间：
//Go对时间和时间段提供了大量的支持
package main

import (
	"fmt"
	"time"
)

func main() {

	//得到当前时间
	now := time.Now()
	fmt.Println(now)

	//通过提供年月日等信息，可以构建一个time
	//时间总是关联着位置信息，例如时区
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237,
		time.UTC)
	fmt.Println(then)

	//可以提取时间的各个组成部分
	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())

	//输出是星期一到日的weekday也是支持的
	fmt.Println(then.Weekday())

	//这些方法来比较两个时间，分别测试一下是否是之前，之后或者是同一时刻，精确到秒
	fmt.Println(then.Before(now))
	fmt.Println(then.After(now))
	fmt.Println(then.Equal(now))

	//方法sub返回一个duration来表示两个时间点的间隔时间
	diff := now.Sub(then) //现在的时间减去之前的时间
	fmt.Println(diff)

	//我们计算出不同单位下的时间长度值
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Nanoseconds())

	//可以使用add将时间后移一个时间间隔
	//或者使用一个'-'来将时间前移一个时间间隔
	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))
}
