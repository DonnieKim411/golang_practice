package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := printSquare(square, 10)

	fmt.Println(x)

}

func square(n int) int {

	return n * n

}

func printSquare(f func(int) int, feed int) string {

	fmt.Println("we are squaring: ", feed)
	squared := f(feed)
	fmt.Println("squard: ", squared)

	return strconv.Itoa(squared)
}
