package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println(p.first, p.last, "is speaking")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Donnie",
		last:  "Kim",
		age:   33,
	}
	var p2 = person{
		first: "Claire",
		last:  "Kim",
		age:   3,
	}

	p1.speak()

	saySomething(&p1)

	// saySomething(p1) This will fail because speak requires a pointer receiver
	// but p1.speak() works as it will figure out. So when working with interface,
	// ensure which receiver it requres, and what value type u are passing in

	// receivers 		values
	// ----------------------
	// (t T)		   	T and *T
	// (t *T)			*T

}
