package main

import "fmt"

// range 是一个内置函数，可以遍历数组、切片slice、映射map、字符串
func range1() {
	// 遍历切片
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	sum := 0
	for i, v := range pow {
		sum += v
		fmt.Println(i, v, sum)
	}
	fmt.Println("sum: ", sum)

	// 遍历映射
	kw := map[string]string{"go": "golang", "py": "python"}
	for k, v := range kw {
		fmt.Printf("%s: %s\n", k, v)
	}

	// 遍历字符串
	for i, c := range "golang" {
		fmt.Printf("%d, %c\n", i, c)
	}
}

// 可以将下标或值赋予 _ 来忽略它。若你只需要索引，去掉 , value 的部分即可。
func range2() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	range1()
	range2()
}
