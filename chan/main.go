package main

import (
	"fmt"
	"sync"
)

var a int

var b chan int // 指定通道中的元素类型

func chanlDemo() {
	fmt.Println(b)
	b = make(chan int) // 不带缓冲区的初始化
	go func() {
		x := <-b
		fmt.Println("通道x:", x)
	}()
	b <- 10 //放进去没放出，死锁
	fmt.Println("10放到通道")
	b = make(chan int, 16) // 带缓冲区初始化
	fmt.Println(b)
}

var wg sync.WaitGroup

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(2)
	go f1(a)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
