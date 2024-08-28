package main

import (
	"fmt"
	"puppy/ex87-BET2/quote"
	"puppy/ex87-BET2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
