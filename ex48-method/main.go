package main

import "fmt"

type Person struct {
	first string
	age   int
}

func (p Person) Speak() {
	fmt.Println(p.first, p.age)
}

func main() {
	p := Person{
		first: "donnie",
		age:   33,
	}
	p.Speak()
}
