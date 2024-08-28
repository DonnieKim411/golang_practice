package main

import (
	"fmt"
)

func main() {

	for y := 0; y < 10; y++ {
		if y%2 == 1 {
			fmt.Println("odd number!")
			continue
		}
		fmt.Println(y)

	}

}
