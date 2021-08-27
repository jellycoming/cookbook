package main

import "fmt"

// 数组由一组数据类型相同的值组成。类型 [n]T 表示拥有 n 个 T 类型的值的数组。
// 数组的长度是其类型的一部分，因此数组不能改变大小。

func array() {
	// 将变量 i 声明为长度为10的 int 类型数组，数组中的每个元素会被自动初始化为零值
	// len 函数求数组的长度，cap 函数求数组的容量
	var i [10]int
	fmt.Println(i)                         // [0 0 0 0 0 0 0 0 0 0]
	fmt.Printf("%d, %d\n", len(i), cap(i)) // 10, 10

	var a [3]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)

	// 声明同时初始化
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func main() {
	array()
}
