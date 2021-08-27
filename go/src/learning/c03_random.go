package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成随机数
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("random int: ", r.Intn(100))

	rand.Seed(time.Now().UnixNano())
	fmt.Println("random int: ", rand.Intn(100))

	// 返回 1 - 10 洗牌后的 []int 切片
	for i, v:= range rand.Perm(10){
		fmt.Printf("%d: %d\n", i, v)
	}
}
