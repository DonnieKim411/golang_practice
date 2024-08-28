package main

import (
	"fmt"
)

func main() {
	q := make(chan int) // bidirectional channel
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		q <- 101
		close(c)

	}()

	return c
}

func receive(c, q <-chan int) {
	for c != nil || q != nil {
		select {
		case v, ok := <-c:
			if ok {
				fmt.Println("from c: ", v)
			} else {
				c = nil // Ensure c is nil so it's not checked again
			}
		case v, ok := <-q:
			if ok {
				fmt.Println("from q:", v)
			}
			q = nil // Ensure q is nil so it's not checked again and loop can terminate
		}
	}
}
