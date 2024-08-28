package main

import (
	"fmt"
)

func main() {
	r1 := []string{"James", "Bond", "Shaken, not stirred"}
	r2 := []string{"Miss", "Moneypenny", "I'm 008."}

	xp := [][]string{r1, r2}
	fmt.Printf("%v - %T\n", xp, xp)
	for i, v := range xp {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}

}

/*
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
*/
