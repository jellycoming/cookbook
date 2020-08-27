package main

import (
	"fmt"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

// 切片是在数组上抽象的一个数据类型，为数组元素提供动态大小的、灵活的视角。
func slice1() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// 通过数组创建切片
	var s1 []int = primes[1:4]
	fmt.Printf("%v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	// 通过make关键字创建切片
	s2 := make([]int, 0)
	fmt.Printf("%v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	// 直接声明切片，文法类似于没有长度的数组。实际上会首先创建一个数组，然后构建一个引用了它的切片。
	var s3 = []int{1, 1, 2, 3, 5, 8, 13, 21, 34}
	fmt.Printf("%v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	s4 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Printf("%v, len=%d, cap=%d\n", s4, len(s4), cap(s4))
}

// 切片并不存储任何数据，它只是描述了底层数组中的一段。切片就像数组的引用。
// 更改切片的元素会修改其底层数组中对应的元素。与它共享底层数组的切片都会观测到这些修改。
// 切片操作符 : 可以忽略上下界。参考 Python
func slice2() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a := names[:2]
	b := names[1:3]
	c := names[:]
	fmt.Println(a, b, c)

	b[0] = "XXX" // 数组names会被修改，同时切片 a 与切片 c 也会观测到这个修改
	fmt.Println(a, b, c)
	fmt.Println(names)
}

// 切片的长度就是它所包含的元素个数。
// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
func slice3() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // [2 3 5 7 11 13], len=6, cap=6

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s) // [], len=0, cap=6

	// 拓展其长度
	s = s[:4]
	printSlice(s) // [2 3 5 7], len=4, cap=6

	// 舍弃前两个值
	s = s[2:]
	printSlice(s) // [5 7], len=2, cap=4

	// 切片的零值是 nil ，nil 切片的长度和容量为 0 且没有底层数组。
	var s0 []int
	printSlice(s0) // [], len=0, cap=0
	if s0 == nil {
		fmt.Println("slice s0 is nil!")
	}
}

// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片。要指定它的容量，需向 make 传入第三个参数。
func slice4() {
	a := make([]int, 5)
	printSlice(a) // [0 0 0 0 0], len=5, cap=5
	b := make([]int, 0, 5)
	printSlice(b) // [], len=0, cap=5
	c := b[:2]
	printSlice(c) // [0 0], len=2, cap=5
	d := c[2:5]
	printSlice(d) // [0 0 0], len=3, cap=3
}

// 切片可包含任何类型，甚至包括其它的切片。
func slice5() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// 可以用 append 关键字为切片追加新元素。
// 当 append 新元素的切片的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。
func slice6() {
	// 定义一个长度为 1 的 int 类型切片
	var s = make([]int, 1)
	printSlice(s)
	fmt.Println(&s[0])
	s[0] = 1
	printSlice(s)
	fmt.Println(&s[0]) // 到这 s[0] 的内存地址还没有变，底层数组还是原来的那个

	s = append(s, 2)   // 再次 append 新元素的时候，原来的切片容量已经不够了，会重新分配一个更大的数组，切片 s 会指向这个新分配的数组。
	fmt.Println(&s[0]) // 到这 s[0] 的内存地址就是重新分配的了
	printSlice(s)
}

func main() {
	slice1()
	slice2()
	slice3()
	slice4()
	slice5()
	slice6()
}
