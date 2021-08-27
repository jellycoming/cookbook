package main

import "fmt"

// 类型断言 提供了访问接口值底层具体值的方式。
// t := i.(T)
// 该语句断言接口值 i 保存了具体类型 T，并将其底层类型为 T 的值赋予变量 t。
func typeof() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	//f = i.(float64) // panic
}

// 类型选择
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

/**
fmt 包中定义的 Stringer 是最普遍的接口之一。

	type Stringer interface {
		String() string
	}

Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。
*/
type Person struct {
	Name string
	Age int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	typeof()
	do(21)
	do("hello")
	do(true)

	p1:= Person{"xx", 25}
	p2 := Person{"oo", 30}
	fmt.Println(p1, p2)
}
