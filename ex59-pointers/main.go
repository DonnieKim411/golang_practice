package main

import "fmt"

type person struct {
	first string
	age   int
}

func main() {
	x := 42
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", &x, &x)

	p1 := person{
		first: "donnie",
		age:   33,
	}

	fmt.Printf("%v %T\n", p1, p1)
	fmt.Printf("%v %T\n", &p1, &p1)
	fmt.Printf("%v %T\n", &p1.first, &p1.first)
	fmt.Printf("%v %T\n", &p1.age, &p1.age)
}
