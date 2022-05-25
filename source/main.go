package main

import "fmt"

func main() {
	var a *int // nil pinter
	//*a = 100

	fmt.Println(a)

	var a1 = new(int)
	fmt.Println(a1)
}

// make 和 new 的区别
// new 很少用，一般给基本数据类型申请内存，string，int，返回是对应类型的指针
// make 是用来给slice、map、chan申请内存，make函数返回的是对应这三个类型本身
