package main

import (
	"fmt"
)

func main() {
	// create a slice of type int
	xi := [10]int{}
	fmt.Printf("%T \t %#v\n", xi, xi)

	for i := 0; i < 10; i++ {
		xi[i] = i + 42
	}
	fmt.Printf("%T \t %v\n", xi, xi)

	for i, v := range xi {
		fmt.Printf("%v - %T - %v\n", v, v, i)

	}
}
