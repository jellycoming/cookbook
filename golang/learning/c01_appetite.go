// 包名，main包为程序入口
package main

// 导入其他包
// 按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。
import (
	"cookbook-go/learning/constant"
	"fmt"
	"math/rand"
)

// 特殊的初始化方法，先于main方法执行
func init() {
	fmt.Println("Initialize")
}

// 程序入口方法
func main() {
	fmt.Println("Hello World", rand.Intn(10))
	// 在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的，在导入一个包时，你只能引用其中已导出的名字
	fmt.Println(constant.ON)
}
