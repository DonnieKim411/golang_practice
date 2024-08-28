package main

import (
	"fmt"
)

// func main() {
// 	// cs := make(chan<- int) this is send channel. we need to make it bidirectional to receieve from the channel
// 	cs := make(chan int)

// 	go func() {
// 		cs <- 42
// 	}()
// 	fmt.Println(<-cs)

// 	fmt.Printf("------\n")
// 	fmt.Printf("cs\t%T\n", cs)
// }

func main() {
	// cr := make(<-chan int) this is a receieve only channel. Again, make it bidirectional so that we can send to the channel
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
