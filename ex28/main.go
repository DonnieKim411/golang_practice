package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"james": 42,
		"money": 199,
	}

	for k, v := range m {
		fmt.Printf("key %v \t value %v\n", k, v)
	}
	fmt.Println("--------")
	age := m["james"]
	fmt.Println(age)

	fmt.Println("--------")
	if age, ok := m["james"]; ok {
		fmt.Println("james age: ", age)
	}
	if _, ok := m["donnie"]; !ok {
		fmt.Println("donnie age does not exist in our map")
	}
	fmt.Println("donnie age: ", m["donnie"])
}
