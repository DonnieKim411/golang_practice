package main

import "fmt"

type person struct {
	first string
}

func changeName(p person, newName string) person {
	p.first = newName
	return p
}

func changeNamePtr(p *person, newName string) {
	p.first = newName

}

func main() {
	p := person{
		first: "donnie",
	}
	fmt.Println(p)

	p = changeName(p, "claire")
	fmt.Println(p)

	changeNamePtr(&p, "ascher")
	fmt.Println(p)
}
