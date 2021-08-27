package main

import (
	"fmt"
	"time"
)

// 没有参数的函数
func say() string {
	return "Hello World!"
}

// 没有返回值的函数
func sayx() {
	fmt.Println("Hi Baby!")
}

// 有多个返回值的函数
func says() (string, int, string) {
	return "Hi sweet", 10000, "times"
}

// 有两个int类型参数的函数
func add(x, y int) int {
	return x + y
}

// 有多个不同类型参数的函数
func adds(x, y int, s string) string {
	return fmt.Sprintf("%s: %d", s, x+y)
}

// 命名返回值的函数
// x,y会被视作定义在函数顶部的变量，没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 没有参数名，只有参数类型的方法。
// 这种写法常用于方法只是为了实现某接口，而具体实现逻辑中并不关心传入的参数。
func noArgsName(int, string) {
	fmt.Sprintln("no args name func")
}

// 参数类型为函数的函数，可以接收普通函数和结构体函数
func fArgs(f func()) {
	f()
}

func hello() {
	fmt.Println("hello world")
}

type Hello struct {
	Timestamp int64
}

func (h *Hello) hello() {
	fmt.Println("hello world", h.Timestamp)
}

func main() {
	fmt.Println(say())
	sayx()
	fmt.Println(says())
	fmt.Println(add(2, 3))
	fmt.Println(adds(2, 3, "The result is"))
	fmt.Println(split(17))
	noArgsName(11, "hello")
	// 调用参数类型为函数的函数
	fArgs(hello)
	h := new(Hello)
	h.Timestamp = time.Now().Unix()
	fArgs(h.hello)
}
