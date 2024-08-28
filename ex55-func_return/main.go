package main

import "fmt"

func main() {
	x := outer()
	fmt.Println(x())
}

func outer() func() int {

	fmt.Println("outer!")

	return func() int {
		return 42
	}

}
