package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// 返回当前时间: 2019-01-17 19:19:51.376874 +0800 CST m=+0.000349499
	t := time.Now()
	fmt.Println(t)
	// 返回当前Unix时间戳: 1547723991
	fmt.Println(time.Now().Unix())
	// 返回当前纳秒级Unix时间戳: 1547725263327011000
	fmt.Println(time.Now().UnixNano())
	// 当前星期
	fmt.Println(t.Weekday())
	// 当前星期字符串表示
	fmt.Println(strconv.Itoa(int(t.Weekday())))
	// 当前小时
	fmt.Println(t.Hour())
	// 将时间转为不同格式的字符串
	fmt.Println(t.Format("2006-01-02")) // 转为"yyyy-mm-dd"格式的字符串: 2019-01-17
	fmt.Println(t.Format("2006-01-02 15:04:05")) // 转为"yyyy-mm-dd HH:MM:SS"格式的字符串: 2019-01-17 19:19:51
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format("3:04PM"))
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t.Format("2006-01-02T15:04:05.999999-07:00"))
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// 将字符串转为不同格式的时间
	t1, e := time.Parse(time.RFC3339, "2019-01-01T19:59:59+00:00")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(t1)
	t2, e := time.Parse("2006-01-02", "2019-01-17")
	fmt.Println(t2)
	t3, e := time.Parse("2006-01-02 15:04:05", "2019-01-17 00:00:00")
	fmt.Println(t3)
}
