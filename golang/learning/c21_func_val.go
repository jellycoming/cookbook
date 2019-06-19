package main

import (
	"fmt"
	"math"
	"reflect"
)

// 在 Go 语言中，函数也是值。它们可以像其它值一样传递。函数值可以用作函数的参数或返回值。
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// 闭包是引用了自由变量的函数。这个被引用的自由变量将和这个函数一同存在，即使已经离开了创造它的环境也不例外。
// Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被“绑定”在了这些变量上。
func closure() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println(reflect.TypeOf(hypot)) // hypot 的类型为：func(float64, float64) float64

	// 闭包
	pos, neg := closure(), closure()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
