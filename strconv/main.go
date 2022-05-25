package main

import (
	"fmt"
	"strconv"
)

//strconv

func main() {

	// 字符串转换为Int
	str := "10000"
	ret1, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(ret1)
	fmt.Printf("%v,%T", ret1, ret1)

	// 数字转换为字符串

}
