package main

import "fmt"

// 接口是引用类型
// 空接口是数据类型，任何数据类型都可以赋予给空接口

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (s Stu) Say() {
	fmt.Println("Stu Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("inter Sqy i = ", i)
}

func main() {
	var stu Stu
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say()
}
