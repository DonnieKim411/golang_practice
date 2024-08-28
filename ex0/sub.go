package main

import "fmt"

var y int

func main() {
	fmt.Println(y)
	z := 42.0
	var foo string = "happy"

	fmt.Printf("%v is type %T\n", z, z)
	fmt.Printf("%v is type %T\n", foo, foo)
}
