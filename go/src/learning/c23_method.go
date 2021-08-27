package main

import (
	"fmt"
	"math"
)

// 方法就是一类带特殊的 接收者 参数的函数。
// 方法接收者可以是结构体类型或者同一包内的自定义类型
type Vertexf struct {
	X float64
	Y float64
}

type MyFloat float64

// 为结构体类型声明方法
func (v Vertexf) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 指针接收者。若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。
// 指针接收者的方法可以修改接收者指向的值，同时避免在每个方法调用时都进行值拷贝，若值的类型为大型结构体时，这样做会更加高效。
func (v *Vertexf) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 带指针的函数
func Scale(v *Vertexf, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 方法只是个带接收者参数的函数。现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。
func Abs(v Vertexf) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 为非结构体类型声明方法。接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 函数调用时参数的指针类型和值类型必须严格匹配
// 方法调用时调用者的类型会根据方法参数的类型被 Go 语言重新解释
func main() {
	v := Vertexf{3, 4}
	fmt.Println(v.Abs()) // 5
	p := &v
	fmt.Println(p.Abs()) // 以值为接收者的方法被调用时，接收者既能为值又能为指针。Go 会将方法调用 p.Abs() 解释为 (*p).Abs()。
	fmt.Println(Abs(v))

	v.Scale(10) // 以指针为接收者的方法被调用时，接收者既能为值又能为指针。Go 会将语句 v.Scale(10) 解释为 (&v).Scale(10)
	fmt.Println(v.Abs()) // 输出变为50。由于是指针接收者，这里改变了v的属性，如果变成值接收者，这里不会改变

	Scale(&v, 10) // 带指针参数的函数调用时必须接受一个指针
	fmt.Println(Abs(v))

	f := MyFloat(math.Sqrt2)
	fmt.Println(f.Abs())
}
