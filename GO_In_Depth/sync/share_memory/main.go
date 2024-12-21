package main

import (
	"fmt"
	"sync"
)


//通过加锁 解锁 f
//通过共享内存 实现并发 进而实现通信
func main() {
	num := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			num++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(num)
}
