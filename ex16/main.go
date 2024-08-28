package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)

	fmt.Printf("rand number is %v\n", x)

	if x > 0 && x < 100 {
		fmt.Println("betwee 0 and 100")
	} else if x >= 100 && x < 200 {
		fmt.Println("betwee 100 and 200")
	} else {
		fmt.Println("betwee 200 and 250")
	}
}
