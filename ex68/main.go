package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Num Go routine:\t", runtime.NumGoroutine())
	fmt.Println("Num CPU:\t", runtime.NumCPU())

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("hello")
		wg.Done()
	}()

	go func() {
		fmt.Println("world")
		wg.Done()
	}()

	fmt.Println("-------------Before wait-------------")
	fmt.Println("Num Go routine:\t", runtime.NumGoroutine())
	fmt.Println("Num CPU:\t", runtime.NumCPU())

	wg.Wait()
	fmt.Println("-------------After wait-------------")
	fmt.Println("Num Go routine:\t", runtime.NumGoroutine())
	fmt.Println("Num CPU:\t", runtime.NumCPU())
}
