package main

import "fmt"

func main() {

	s1 := struct {
		first      string
		friends    map[string]int
		fav_drinks []string
	}{
		first: "donnie",
		friends: map[string]int{
			"claire":   3,
			"duckwhan": 36,
			"minjung":  39,
		},
		fav_drinks: []string{
			"Coke Zero",
			"Diet Dr Pepper",
		},
	}

	for _, v := range s1.friends {
		fmt.Println(v)
	}

	for _, v := range s1.fav_drinks {
		fmt.Println(v)
	}

}
