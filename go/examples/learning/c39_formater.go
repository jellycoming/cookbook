package main

import "fmt"

type Campaign struct {
	Code    int
	Name    string
	Price   float64
	Serving bool
}

var campaign = Campaign{
	Code:    101,
	Name:    "campaign for my game",
	Price:   50.50,
	Serving: true,
}

func main() {
	// 通用占位符
	// %v 值的默认格式
	fmt.Printf("%v\n", campaign)         // {1001 campaign for my game 50.5 true}
	fmt.Printf("%v\n", campaign.Price)   // 50.5
	fmt.Printf("%v\n", campaign.Serving) // true
	// %+v 打印结构体时，会添加字段名
	fmt.Printf("%+v\n", campaign) // {Code:1001 Name:campaign for my game Price:50.5 Serving:true}
	// %#v 值的Go语法表示，会为字符串加上双引号 ""
	fmt.Printf("%#v\n", campaign) // main.Campaign{Code:1001, Name:"campaign for my game", Price:50.5, Serving:true}
	// %T 值得类型的Go语法表示
	fmt.Printf("%T\n", campaign) // main.Campaign

	// 布尔型占位符
	fmt.Printf("%t\n", campaign.Serving) // true

	// 整数占位符
	// 二进制
	fmt.Printf("%b\n", campaign.Code) // 1100101
	// 十进制
	fmt.Printf("%d\n", campaign.Code) // 101
	// 八进制
	fmt.Printf("%o\n", campaign.Code) // 145
	// 十六进制
	fmt.Printf("%x\n", campaign.Code) // 65
	// 十六进制
	fmt.Printf("%X\n", campaign.Code) // 65

	// 浮点数占位符
	// 默认 6 位小数
	fmt.Printf("%f\n", campaign.Price) // 50.500000
	// 保留 4 位小数
	fmt.Printf("%.4f\n", campaign.Price) // 50.5000
	// 总共 6 位，小数位数为 3 位
	fmt.Printf("%6.3f\n", campaign.Price) // 50.500
	// 忽略小数末尾的 0
	fmt.Printf("%g\n", campaign.Price) // 50.5
	// 忽略小数末尾的 0
	fmt.Printf("%G\n", campaign.Price) // 50.5
	// 科学计数法
	fmt.Printf("%e\n", campaign.Price) // 5.050000e+01
	fmt.Printf("%E\n", campaign.Price) // 5.050000E+01

	// 字符串占位符
	fmt.Printf("%s\n", campaign.Name) // campaign for my game
	fmt.Printf("%q\n", campaign.Name) // "campaign for my game"

	// 指针占位符
	fmt.Printf("%p\n", &campaign) // 0x11607a0
}
