package main

import "fmt"

func main() {
	fmt.Println(foo())
	i, s := bar()
	fmt.Println(i, s)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "hello"
}
