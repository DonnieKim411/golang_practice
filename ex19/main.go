package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("this is from init func")
}

func main() {

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("rand number x is %v\n", x)
	fmt.Printf("rand number y is %v\n", y)

	// if x < 4 && y < 4 {
	// 	fmt.Println("both x and y are less than 4")
	// } else if x > 6 && y > 6 {
	// 	fmt.Println("both x and y are greater than 6")
	// } else if x >= 4 && x <= 6 {
	// 	fmt.Println("4 <= x <= 6")
	// } else if y != 5 {
	// 	fmt.Println("y is not 5")
	// } else {
	// 	fmt.Println("none of the conds are met")
	// }
	switch {
	case x < 4 && y < 4:
		fmt.Println("both x and y are less than 4")
	case x > 6 && y > 6:
		fmt.Println("both x and y are greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("4 <= x <= 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("none of the conds are met")
	}
	// switch {
	// case x > 0 && x < 100:
	// 	fmt.Println("betwee 0 and 100")
	// case x >= 100 && x < 200:
	// 	fmt.Println("betwee 100 and 200")
	// case x >= 200:
	// 	fmt.Println("betwee 200 and 250")
	// }

}
