package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH:", runtime.GOARCH)
	fmt.Println("Goroutines:", runtime.NumGoroutine())

}
