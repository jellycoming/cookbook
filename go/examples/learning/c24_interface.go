package main

import (
	"fmt"
	"math"
)

// 接口类型 是由一组方法签名定义的集合。
// 接口类型的变量可以保存任何实现了这些方法的值。
// 类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。
// 隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。
type Abser interface {
	Abs() float64
}

type Vertexi struct {
	X, Y float64
}

type MyFloati float64

func (v *Vertexi) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertexi) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (f MyFloati) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func desc(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var a Abser
	f := MyFloati(-math.Sqrt2)
	v := Vertexi{3, 4}
	a = f // MyFloat 实现了 Abser
	fmt.Println(a.Abs())
	a = &v // *Vertex 实现了 Abser
	fmt.Println(a.Abs())
	//a = v // 编译错误。v 是一个 Vertex（而不是 *Vertex），所以没有实现 Abser。

	// 接口值可以看做包含值和具体类型的元组：(value, type)。接口值调用方法时会执行其底层类型的同名方法。
	var i I
	i = &T{S: "Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	// 即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。保存了 nil 具体值的接口其自身并不为 nil。
	var t *T
	i = t
	describe(i)
	//i.M() // 这里会引发 nil pointer 的运行时错误

	// nil 接口值既不保存值也不保存具体类型。
	var ii I
	describe(ii)
	//ii.M() // 这里会引发 nil pointer 的运行时错误

	// 空接口。空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
	// 空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
	var none interface{}
	desc(none)

	none = 42
	desc(none)

	none = "Hello"
	desc(none)
}
