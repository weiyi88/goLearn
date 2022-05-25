package main

import "fmt"

func funcA() {
	fmt.Println("a")
}

func funcB() {
	defer func() {
		err := recover() // 不推荐，只能使用在特定场所
		// 如果程序出现panic 错误，可以通过recover恢复过来
		if err != nil {
			fmt.Println("做出recover 补救")
		}
		fmt.Println("做错补救")
	}()
	panic("出现了严重错误") // 程序崩溃退出
	//fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
