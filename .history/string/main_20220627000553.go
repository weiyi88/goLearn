package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n = 100

	fmt.Printf("%T\n", n)

	s := "aring"
	fmt.Println(s)

	c := `
人面桃花
人比桃花瘦
`
	word := "here"
	//fmt.Println(c)
	aa := fmt.Sprintf("%s%s", c, word)
	//fmt.Println(aa)

	//fmt.Println(strings.Contains(aa, "人面"))
	//fmt.Println(strings.Index(aa, "瘦"))

	fmt.Println("----------------")
	for _, c := range aa {
		//fmt.Println(c, a)
		fmt.Printf("%c\n", c)
	}

	// unit8 类型，或者byte类型，代表ASCII码的一个自渎
	// rune 类型，代表一个UTF-8字符

	// 字符串底层是一个byte数组，所以可以和[]byte类型互相转换，字符串是不能修改的，字符串是由byte字节组成
	// 所以字符串的长度是byte字节的长度，rune类型用来表示utf8字符，一个rune字符由一个或者多一个byte组成

	fmt.Println("----------------")
	//修改字符串

	s1 := "这是一条字符串"
	s2 := []rune(s1) // 字符串强制转换成一个rune切片
	s2[0] = '新'
	fmt.Println(string(s2))

	fmt.Println("---------------------------")

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
}
