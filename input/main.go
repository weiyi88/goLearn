package main

import (
	"bufio"
	"fmt"
	"os"
)

func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin) // 这样读，才能读出空格后读字符
	s, _ = reader.ReadString('\n')      // 直到读到你换行
	fmt.Printf("你输入的内容是：%s\n", s)
}

func main() {
	useBufio()
}
