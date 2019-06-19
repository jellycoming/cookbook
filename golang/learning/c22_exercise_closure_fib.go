package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	x1, x2 := 0, 1
	return func() int {
		x1, x2 = x2, x1+x2
		return x1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", f())
	}
}
