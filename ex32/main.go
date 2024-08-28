package main

import (
	"fmt"
)

func main() {
	// create a slice of type int
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// fmt.Printf("%T \t %#v\n len(xi)", xi, xi)

	fmt.Printf("%v\n", xi[:5])
	fmt.Printf("%v\n", xi[5:])
	fmt.Printf("%v\n", xi[2:7])
	fmt.Printf("%v\n", xi[3:8])
}
