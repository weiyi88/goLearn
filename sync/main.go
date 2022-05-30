package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex // 读写互斥锁

func add() {
	for i := 0; i < 500000; i++ {
		lock.Lock() // 互斥锁
		x += 1
		lock.Unlock()
		fmt.Println(x)
	}
	wg.Done()
}

// 写锁就是阻塞
func write() {
	defer wg.Done()
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock()
}

// 读锁，上锁后，其他还可以上读锁
func read() {
	defer wg.Done()
	rwlock.RLock() // 加读锁
	fmt.Println(x)
	time.Sleep(10 * time.Millisecond)
	rwlock.RUnlock() // 解读锁
}

func main() {
	start := time.Now()

	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
