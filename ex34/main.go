package main

import (
	"fmt"
)

func main() {
	// create a slice of type int
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	xi = append(xi[:3], xi[6:]...)
	fmt.Println(xi)

}
