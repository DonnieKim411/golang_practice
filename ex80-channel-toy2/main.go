// SOLUTION 1
// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	x := 10
// 	y := 10
// 	c := gen(x, y)

// 	for i := 0; i < x*y; i++ {
// 		fmt.Println(i, <-c)
// 	}
// 	fmt.Println("ROUTINES", runtime.NumGoroutine())
// }

// func gen(x, y int) <-chan int {
// 	c := make(chan int)

// 	for i := 0; i < x; i++ {
// 		go func() {
// 			for j := 0; j < y; j++ {
// 				c <- j
// 			}
// 		}()
// 		fmt.Println("ROUTINES", runtime.NumGoroutine())
// 	}
// 	return c
// }

/*
	--------------------------------------------------------------------------

SOLUTION 2
*/
package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(m int) {
				for i := 0; i < 10; i++ {
					c <- i*10 + m
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
