package main

import "fmt"

func main() {
	// 1, 自定义切片
	var s1 []int // 没定义长度就是切片
	var s2 []string
	var s3 [4]int
	fmt.Println(s3)

	s2 = []string{"aring", "jack", "luck"}

	fmt.Println(s1)
	fmt.Println(s2, cap(s2), len(s2))

	s2 = []string{"chacha", "hahha"}
	fmt.Println(s2, cap(s2), len(s2))

	fmt.Println("-----------------")

	// 由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	a2 := a1[0:4]
	fmt.Println(a2)

	a3 := make([]int, 2, 3)
	a3[0] = 88
	a3[1] = 99
	a3 = append(a3, 4, 34, 34)
	fmt.Println(a3)
	fmt.Println("-------------")

	a4 := []int{1, 3, 5, 7, 8, 9}
	a5 := a4
	a5[2] = 55
	fmt.Println(a4, a5, cap(a5), len(a5))

	// append 接收必须有用切片接收
	// 不然就赋值新的变量
	a5 = append(a5, 77, 88, 99, 10, 11, 12, 123)
	a6 := append(a5, 77, 88, 99, 10, 11, 12, 123)
	fmt.Println(a6, cap(a6), len(a6))
	fmt.Println(a5, cap(a5), len(a5))

	fmt.Println("--------------")
	//var a7 []int	// 都是nil 没有空间
	a7 := make([]int, 5, 9)
	copy(a7, a5)
	a7[4] = 990
	fmt.Println(a7, a5)

	//切片删除
	fmt.Println("--------------------")
	a7 = append(a7[:1], a7[2:]...) //将索引2 的数据删除
	a7 = append(a7[:1], a7[:]...)
	fmt.Println(a7)

}

// 切片只想一个底层数组
// 切片长度就是它元素的个数
// 切片容量就是底层数组从切片第一个元素到最后一个元素到数量
// 切片是引用类型
