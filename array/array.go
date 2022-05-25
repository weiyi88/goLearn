package main

import "fmt"

func main() {

	// 数组
	// 存放元素的容器
	// 必须制定存放的元素类型和容量（长度）
	// 数组的长度是数组类型的一部分

	// 数组不初始化，默认都是0值
	var a1 [3]int
	var a2 [4]bool

	fmt.Println(a1)
	fmt.Println(a2)

	citys := [...]string{"北京", "广州", "深证"}

	// a index
	// v 值
	for a, v := range citys {
		fmt.Println(a, v)
	}

	fmt.Println("---------------------")
	a3 := [...]int{1, 3, 5, 7, 8}
	result := 0

	for _, v := range a3 {
		result += v
	}
	fmt.Println(result)
}
