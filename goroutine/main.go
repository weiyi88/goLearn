package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Wg sync.WaitGroup

func chuanxing() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}
	for i := 20; i < 30; i++ {
		fmt.Println(i)
	}
}

func tryGoroutine() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 20; i < 30; i++ {
			fmt.Println(i)
		}
	}()

	time.Sleep(time.Second)
}

func fi(i int) {
	defer Wg.Done() // 函数结束就，goroutine -1
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

func main() {
	// 开启10个goroutine
	Wg.Add(10)
	for i := 0; i <= 10; i++ {
		go fi(i)
	}
	Wg.Wait() // 等待所有协程池完结,计数器减为0

}
