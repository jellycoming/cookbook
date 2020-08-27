package main

import "fmt"

// 定义一个结构体，结构体及其属性的命名都为首字母大写，表示该结构体及其属性是可导出的，可供外部调用
type Vertex struct {
	X int
	Y int
	M map[string]string
}

// 结构体文法通过直接列出字段的值来新分配一个结构体。
var (
	v1 = Vertex{1, 2, make(map[string]string)} // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}                          // 如果只初始化部分属性，必须指定属性名，Y:0 被隐式地赋予
	v3 = Vertex{}                              // X:0 Y:0
	pp = &Vertex{X: 1, Y: 2}                   // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	// 初始化结构体
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v) // {1 2}

	// 访问及修改结构体中属性 X 的值
	fmt.Println(v.X) // 1
	v.X = 10
	fmt.Println(v) // {10 2}

	// 用指针访问结构体，这里不需要 (*p).X 这种写法，而是直接使用 p.X 这种隐式间接引用
	p := &v
	p.X = 100
	fmt.Println(v) // {100 2}

	fmt.Println(v1, v2, v3, pp)

	// 初始化结构体时，其中的属性如果未指定，会在结构体初始化时被初始化为初始值
	d := Vertex{}
	fmt.Printf("default X: %d, Y: %d, M: %v", d.X, d.Y, d.M) // default X: 0, Y: 0, M: map[]
	for k, v := range d.M {
		fmt.Printf("default %s=%s", k, v) // not reach
	}
}
