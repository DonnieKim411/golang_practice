package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	incrementor := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			x := incrementor
			runtime.Gosched()
			x++
			incrementor = x
			fmt.Println(incrementor)
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("End val:\t", incrementor)
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}
