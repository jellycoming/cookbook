package main

import "fmt"

/**
const关键字用来声明常量。
常量可以是字符、字符串、布尔值或数值。
常量不能用 := 语法声明。
未指定类型的常量由上下文来决定其类型。
*/

const Pi = 3.14
const (
	// 将 1 左移 100 位来创建一个非常大的数字，即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 1.1
}

func main() {
	const World = "世界"
	fmt.Printf("Hello %v(%T)\n", World, World)
	fmt.Printf("Happy %v(%T) Day\n", Pi, Pi)

	const Truth = true
	fmt.Printf("Go rules?%v(%T)\n", Truth, Truth)

	//  根据上下文来决定 Small 和 Big 是 int 还是 float64
	// fmt.Println(Big) // constant 1267650600228229401496703205376 overflows int
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
