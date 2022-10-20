// @program:     Go_example
// @file:        switch.go
// @create:      2022-10-09 16:17
// @description:

//switch是多分支情况时快捷的条件语句
package main

import (
	"fmt"
	"time"
)

func main() {
	//一个基本的switch
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//在一个case语句中，你可以使用逗号来分割多个表达式
	//在这个例子中，我们还使用了可选的default分支
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//不带表达式的switch是实现if/else逻辑的另一种方式
	//case表达式也可以不使用常量
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//类型开关(type switch)比较类型而非值。可以用来发现一个接口值的类型。
	//在这个例子中，变量t在每个分支中会有相应的类型
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t) //%T打印变量类型
			//fmt.Println(t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
