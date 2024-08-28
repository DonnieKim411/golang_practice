package main

import (
	"fmt"
)

func main() {

	for y := 0; y < 5; y++ {

		for z := 0; z < 5; z++ {
			fmt.Printf("outer loop %v\t inner loop %v\n", y, z)
		}
		fmt.Println("---------")
	}

}
