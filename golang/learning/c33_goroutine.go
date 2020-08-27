package main

import (
	"fmt"
	"time"
)

/**
Go 程（goroutine）是由 Go 运行时管理的轻量级线程。
Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。
*/

func call(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// go 关键字会启动一个新的 Go 程并行执行
	go call("Hello")
	call("World")
}
