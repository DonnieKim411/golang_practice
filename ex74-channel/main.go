package main

import (
	"fmt"
)

// deadlock code!
// func main() {
// 	c := make(chan int)

// 	c <- 42

// 	fmt.Println(<-c)
// }

// soln 1: anonymous func. Preferred method
func main() {
	c := make(chan int)
	go func() {

		c <- 42

	}()

	fmt.Println(<-c)
}

// soln 2: buffered channel
// func main() {
// 	c := make(chan int, 1)

// 	c <- 42

// 	fmt.Println(<-c)
// }
