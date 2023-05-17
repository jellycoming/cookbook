package main

import (
	"fmt"
	"math/cmplx"
)

/**
Go语言的基本类型有:
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // uint8 的别名
rune // int32 的别名, 表示一个 Unicode 码点
float32 float64
complex64 complex128

int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	cmp    complex128 = cmplx.Sqrt(-5 + 12i)
)

/**
没有明确初始值的变量声明会被赋予它们的零值。
数值类型: 0
布尔类型: false
字符串:  ""（空字符串）
*/
var (
	i int     // 0
	f float64 // 0
	b bool    // false
	s string  // ""
)

// 类型转换: 表达式 T(v) 将值 v 转换为类型 T。
func cast() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)
	ii := 42
	ff := float64(ii)
	uu := uint(ff)
	fmt.Println(ii, ff, uu)
}

// 在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。
func typeInference() {
	// 当右边包含未指明类型的数值常量时，变量的类型由等号右侧的值（第一次赋值）推导得出。
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	// int 当右值声明了类型时，新变量的类型与其相同
	ii := i // int
	fmt.Printf("i is %T(%v)\n", i, i)
	fmt.Printf("f is %T(%v)\n", f, f)
	fmt.Printf("g is %T(%v)\n", g, g)
	fmt.Printf("ii is %T(%v)\n", ii, ii)
}

func main() {
	// Printf在输出变量的时候对变量进行了格式化。%T: 值的类型的Go语法表示，%v: 值的默认格式，%q: 双引号围绕的字符串，由Go语法安全地转义。
	const format string = "Type: %T Value: %v\n"
	fmt.Printf(format, ToBe, ToBe)
	fmt.Printf(format, MaxInt, MaxInt)
	fmt.Printf(format, cmp, cmp)
	// 零值
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	cast()
	typeInference()
}
