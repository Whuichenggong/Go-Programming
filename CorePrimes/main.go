package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func singleCorePrimes(numbers []int) []int {
	var primes []int
	for _, n := range numbers {
		if isPrime(n) {
			primes = append(primes, n)
		}
	}
	return primes
}

func multiCorePrimes(numbers []int) []int {
	var primes []int
	var mu sync.Mutex
	var wg sync.WaitGroup

	numWorkers := 4
	chunKSize := len(numbers) / numWorkers

	for i := 0; i < numWorkers; i++ {

		//wg.Add(1) 增加任务计数
		wg.Add(1)
		go func(i int) {
			//wg.Wait() 会阻塞，直到所有 Goroutine 都调用 Done。
			defer wg.Done()
			start := i * chunKSize
			end := (i + 1) * chunKSize
			if i == numWorkers-1 {
				end = len(numbers)
			}

			for _, n := range numbers[start:end] {
				if isPrime(n) {
					mu.Lock()
					primes = append(primes, n)
					mu.Unlock()
				}
			}
		}(i)
	}
	wg.Wait()
	return primes
}

func main() {
	numbers := make([]int, 100000000)
	for i := 1; i <= 100000000; i++ {
		numbers[i-1] = i
	}

	start := time.Now()
	singleCorePrimes(numbers)
	singleCoreDuration := time.Since(start)

	start = time.Now()
	multiCorePrimes(numbers)
	multiCoreDuration := time.Since(start)

	fmt.Println("Single Core Duration:", singleCoreDuration)
	fmt.Println("Multi Core Duration:", multiCoreDuration)

}
