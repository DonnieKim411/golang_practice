package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("this is from init func")
}

func main() {
	x := rand.Intn(250)

	fmt.Printf("rand number is %v\n", x)
	switch {
	case x > 0 && x < 100:
		fmt.Println("betwee 0 and 100")
	case x >= 100 && x < 200:
		fmt.Println("betwee 100 and 200")
	case x >= 200:
		fmt.Println("betwee 200 and 250")
	}

}
