package main

import (
	"fmt"
	"sync"
)
/*
从内存中读取 num 的值。
将值加 1。
将新值写回到内存中。
在并发情况下，不同 Goroutine 的操作步骤可能交叉执行，导致多个 Goroutine 读取到相同的 num 值，并覆盖写回，最终一些自增操作被“丢失”。

为了解决这个问题，我们可以使用 sync.Mutex 来保证同一时刻只有一个 Goroutine 访问 num 变量。

sync.Mutex 是一个互斥锁，它可以保证同一时刻只有一个 Goroutine 访问某个变量。

sync/atomic 包，可以进行原子操作来避免数据竞争：

*/
//本段代码 中的 num++ 是非原子操作，导致多个 Goroutine 并发访问时出现数据竞争。
func main() {
	num := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num++
		}()
	}

	wg.Wait()
	fmt.Println(num)
}
