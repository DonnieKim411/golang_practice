package main

import (
	"fmt"
)

func main() {
	xi := [5]int{}

	for i := 0; i < 5; i++ {
		xi[i] = i
	}

	for i, v := range xi {
		fmt.Printf("%v - %T - %v\n", v, v, i)

	}
}
