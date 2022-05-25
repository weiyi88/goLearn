package main

import (
	"bufio"
	"fmt"
	"os"
)

func useSxan() {
	var s string
	fmt.Print("请输入内容")
	fmt.Scanln(&s) // 输入有空格就拿不到后面的值
	fmt.Printf("你输入的内容是：%s\n", s)
}

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是：%s\n", s)
}

func main() {
	useSxan()
}
