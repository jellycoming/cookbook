package main

import "fmt"

/**
Go 只有一种循环结构：for 循环。
基本的 for 循环由三部分组成，它们用分号隔开：

初始化语句：在第一次迭代前执行
条件表达式：在每次迭代前求值
后置语句：在每次迭代的结尾执行

初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。
一旦条件表达式的布尔值为 false，循环迭代就会终止。
*/

func loop1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 初始化语句和后置语句是可选的。如果只剩下条件表达式了，那么那两个分号也是可以省略的，此时相当于C语言中的 while。
func loop2() {
	sum := 1
	//for ; sum < 1000; {
	//	sum += sum
	//}
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// 死循环
func loop3() {
	i := 0
	for {
		i++
		fmt.Println(i)
		if i > 3 {
			break
		}
	}
}

func main() {
	loop1()
	loop2()
	loop3()
}
