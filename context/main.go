package main

//
//import (
//	"context"
//	"fmt"
//	"sync"
//	"time"
//)
//
//var wg sync.WaitGroup
//
//// 两种通知协程退出
//var notify bool //全局变量
//var exitChan = make(chan bool, 1)
//
//// 为什么需要context
//func f(ctx context.Context) {
//OUT:
//	for {
//		fmt.Println("Aring")
//		time.Sleep(time.Millisecond * 500)
//		select {
//		case <-ctx.Done():
//			break OUT
//		default:
//
//		}
//	}
//	defer wg.Done()
//}
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	wg.Add(1)
//	go f()
//	time.Sleep(time.Second * 5)
//	notify = true
//	exitChan <- true
//	wg.Wait()
//}
