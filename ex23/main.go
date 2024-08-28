package main

import (
	"fmt"
)

func main() {
	y := 0
	for y < 10 {
		if y == 5 {
			fmt.Println("breaking")
			break
		}
		fmt.Println(y)
		y++
	}

}
