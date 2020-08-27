package util

import "fmt"

var Pi = 3.14

/**
包内的同一个源文件也可以有多个 init 函数
*/
func init() {
	fmt.Println("1 init pkg constant, file util.go")
}

func init() {
	fmt.Println("2 init pkg constant, file util.go")
}
