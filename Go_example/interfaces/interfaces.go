// @program:     Go_example
// @file:        interfaces.go
// @create:      2022-10-11 15:54
// @description:

//接口(interfaces)是命名了的方法签名(signatures)的集合
package main

import (
	"fmt"
	"math"
)

//这是一个几何体的基本接口
type geometry interface {
	area() float64
	perim() float64
}

//在例子中，我们将在类型rect和circle上实现这个接口
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

//要在Go中实现一个接口，我们就要实现接口中的所有方法。
//这里我们在rect上实现了geometry接口
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//circle的实现
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//如果一个变量具有接口类型，那么我们可以调用指定接口中的方法
//这里有一个通用的measure函数，利用它来在任何geometry上工作
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	//结构体类型circle和rect都实现了geometry接口
	//所以我们可以使用他们的实例作为measure的参数
	measure(r)
	measure(c)
}
