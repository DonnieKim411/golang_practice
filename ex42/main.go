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
		last:    "Kim2",
		flavors: []string{"chocolate", "strawberry"},
	}

	m1 := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(m1)

	for _, v := range m1 {
		fmt.Println(v)

		for _, v2 := range v.flavors {
			fmt.Println(v.first, v.last, v2)
		}
	}

}
