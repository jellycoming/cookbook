package main

import (
	"fmt"
	"go-exercise/learning/constant"
	"go-exercise/learning/util"
)
/**
1. init 函数用于程序执行前做包的初始化工作
2. 同一个包内可以有多个 init 函数 (参考 constant 包)
3. 包内的同一个源文件也可以有多个 init 函数 (参考 util 包)
4. 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
5. init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
 */
func init() {
	fmt.Println("init pkg main.")
}

func main() {
	fmt.Println(util.Pi)
	fmt.Println(constant.ON)
}
