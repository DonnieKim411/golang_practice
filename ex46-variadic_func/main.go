package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))

}

func foo(i ...int) int {
	fmt.Printf("%T\n", i)
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func bar(i []int) int {
	fmt.Printf("%T\n", i)
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
