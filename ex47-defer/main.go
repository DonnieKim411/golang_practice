package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6}
	defer fmt.Println("this is foo", foo(xi...))
	fmt.Println("this is bar", bar(xi))

}

func foo(i ...int) int {
	fmt.Printf("inside foo %T\n", i)
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func bar(i []int) int {
	fmt.Printf("inside bar %T\n", i)
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
