package main

import (
	"fmt"
	"runtime"
	"time"
)

// 同if语句习惯一样，可以在条件表达式前执行一个简单的语句(os := runtime.GOOS 运行时获取当前的操作系统)。 该语句声明的变量作用域仅在 switch 之内。
// Go语言的switch中除非以 fallthrough 语句结束，否则分支会自动终止。
// Go语言的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。
// switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
func switch1() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		// fallthrough // 如果加上该关键字，下一个紧邻的分支不管条件是否成立也会被执行
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

// 判断距离周六还有多久
func switch2() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// 没有条件的 switch 同 switch true 一样。每一个 case 选项都是 bool 表达式，值为 true 的分支就是被执行的分支。或者执行 default 。
func switch3() {
	t := time.Now()
	fmt.Printf("Now time is: %s\n", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func main() {
	switch1()
	switch2()
	switch3()
}
