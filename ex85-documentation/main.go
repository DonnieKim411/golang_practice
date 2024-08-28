package main

import (
	"ex85-documentation/dog"
	"fmt"
)

type doggo struct {
	name string
	age  int
}

func main() {
	my_dog := doggo{
		name: "kami",
		age:  dog.Years(10),
	}

	fmt.Println("my dog age is:", my_dog.age)

}
