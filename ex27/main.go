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

}
