package main

import "fmt"

// 变量可以定义在包内，也可以定义在函数内
var c, python, java bool
// 变量可在声明时初始化，如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
var x, y int = 1, 2
var z = "3"
// 函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
// err := "err"

func main() {
	var i int
	var j int = 10
	// 在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
	m, n := 0, 0
	// 变量作用域: 在函数内和函数外，有同名的变量时，函数内使用的函数内声明的这个变量。
	x := 11
	fmt.Println(i, j, m, n, c, python, java, x, y, z)
}
