package main

import (
	"fmt"
	"runtime"
	"time"
)

// func a(wg *sync.WaitGroup) {
// 	for i := 1; i < 10; i++ {
// 		fmt.Println("A:", i)
// 	}
// 	wg.Done() // 表示 a 执行完成
// }

// func b() {
// 	for i := 1; i < 10; i++ {
// 		fmt.Println("B:", i)
// 	}
// }

// func main() {
// 	runtime.GOMAXPROCS(1)
// 	var wg sync.WaitGroup
// 	wg.Add(1) // 等待 1 个 goroutine
// 	go a(&wg)
// 	wg.Wait() // 等待 a 执行完
// 	go b()
// 	time.Sleep(time.Second)
// }

//2.

func a(ch chan struct{}) {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
	ch <- struct{}{} // 发送信号，表示 a 执行完成
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	ch := make(chan struct{})
	go a(ch)
	<-ch // 等待 a 执行完
	go b()
	time.Sleep(time.Second)
}
