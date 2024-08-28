package main

import (
	"fmt"
)

func main() {
	c := gen()
	// fmt.Println("Done generating. Now receiving")
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	// fmt.Println("start generating")

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	// fmt.Printf("%T %v\n", c, c)

	for v := range c {
		fmt.Println("Receiving from the channel", v)
	}

}

// USING BUFFERED CHANNEL
// func main() {
// 	c := gen()
// 	fmt.Println("Done generating. Now receiving")
// 	receive(c)

// 	fmt.Println("about to exit")
// }

// func gen() <-chan int {
// 	c := make(chan int, 100)
// 	fmt.Println("start generating")

// 	for i := 0; i < 100; i++ {
// 		c <- i
// 	}
// 	close(c)

// 	return c
// }

// func receive(c <-chan int) {

// 	for v := range c {
// 		fmt.Println("Receiving from the channel", v)
// 	}

// }
