package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	const maxCount = 3 // 定义运行次数
	var wg sync.WaitGroup
	wg.Add(1) // 新 goroutine 计数

	// 新 goroutine
	go func() {
		defer wg.Done() // goroutine 执行结束时通知 WaitGroup
		for i := 1; i <= maxCount; i++ {
			fmt.Printf("new goroutine: i = %d\n", i)
			runtime.Gosched() // 主动让出 CPU
		}
	}()

	// 主 goroutine
	for i := 1; i <= maxCount; i++ {
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		//runtime.Gosched() // 主动让出 CPU
	}

	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("All goroutines finished.")
}
