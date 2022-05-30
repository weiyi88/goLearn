package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		select { // 哪个通道有值就取哪个，随机取
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
