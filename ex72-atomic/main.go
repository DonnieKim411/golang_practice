package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	var incrementor int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)
			fmt.Println(atomic.LoadInt64(&incrementor))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("End val:\t", incrementor)
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}
