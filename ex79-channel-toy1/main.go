package main

import "fmt"

func main() {

	c := make(chan int)

	// go func() {

	// }()
	c = send(c)

	receive(c)

}

func send(c chan int) chan int {

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c

}

func receive(c chan int) {

	for v := range c {
		fmt.Println("Receieved: ", v)
	}

}
