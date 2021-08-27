package main

import "fmt"

// 指针保存变量的内存地址

func pointer() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值（42）
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值（21）

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值（73）
}

func main() {
	// 声明一个指向int类型值的指针p。
	var p *int
	// & 操作符是取地址符，&i 代表的是变量 i 的内存地址。
	i := 42
	p = &i
	fmt.Println(p) // 真实内存地址
	// * 操作符表示指针指向的底层值。
	fmt.Println(*p) // 通过指针 p 读取 i
	*p = 21         // 通过指针 p 设置 i.
	fmt.Println(*p) // 21
	fmt.Println(i)  // 21
	fmt.Println("--------")
	pointer()
}
