package main

import "fmt"

type person struct {
	first, last string
	flavors     []string
}

func main() {

	p1 := person{
		first:   "Donnie",
		last:    "Kim",
		flavors: []string{"chocolate", "vanilla"},
	}

	p2 := person{
		first:   "Claire",
		last:    "Kim",
		flavors: []string{"chocolate", "strawberry"},
	}

	fmt.Println(p1, p1.first, p1.last, p1.flavors)
	fmt.Println(p2, p2.first, p2.last, p2.flavors)

	for i, v := range p1.flavors {
		fmt.Println(i, v)
	}
}
