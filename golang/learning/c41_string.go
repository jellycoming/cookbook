package main

import "fmt"

// 用 + 号进行字符串链接
func str1() {
	fmt.Println("Hello" + " " + "World")
}

func str2() {
	s := "abced"
	fmt.Println(s[0:])
	fmt.Println(len(s))
}

func main() {
	str1()
	str2()
}
