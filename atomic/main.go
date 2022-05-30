package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var l sync.Mutex
var wg sync.WaitGroup

func atomicAdd() {
	defer wg.Done()
	atomic.AddInt64(&x, 1)
}

func main() {
	for i := 0; i < 5000; i++ {
		go atomicAdd()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(x)
}
