package main

import (
	"fmt"
	"time"
)

func TimeFunc() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano()) // 纳秒

	// 接收时间戳
	ret := time.Unix(1652008969, 0)
	fmt.Println(ret)
}

func main() {
	TimeFunc()
}
