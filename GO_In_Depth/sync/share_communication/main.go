package main

import (
	"fmt"
	"sync"
)

/*
因为对 num 的操作完全封装在 Goroutine 内部，外部只能通过 addChan 和 getChan 进行通信，避免了数据竞争。这使用了通信来实现共享内存
*/


type Num struct {
	addChan chan int
	getChan chan chan int
}

func NewNum() *Num {
	addChan := make(chan int)
	getChan := make(chan chan int)

	// go func() { ... }() 表示启动了一个匿名 Goroutine
	go func() {
		num := 0

		//for { ... } 是一个死循环，确保 Goroutine 持续运行。 Goroutine 会一直监听 addChan 和 getChan，等待外界发送的请求
		for {
			//select 是 Go 中用于多路复用的语法，可以同时监听多个 channel
			select {
			case val := <-addChan:
				num += val
			case respChan := <-getChan:
				respChan <- num
			}
		}
	}()
	return &Num{
		addChan: addChan,
		getChan: getChan,
	}
}

func (n *Num) Add(val int) {
	n.addChan <- val
}

func (n *Num) Get() int {
	resChan := make(chan int)
	n.getChan <- resChan
	val := <-resChan
	return val
}

func main() {
	num := NewNum()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			num.Add(1)
		}()

	}
	wg.Wait()
	fmt.Println(num.Get())

}
