package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	v1 := fmt.Sprint("toJSON failed!: ", "sd")
	v2 := fmt.Sprintf("toJSON failed!: %v", p1.Sayings[0])
	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Printf("%T\n", v1)

	bs, err := toJSON(p1)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		// return []byte{}, fmt.Errorf("toJSON failed!: %v", err)
		return []byte{}, errors.New(fmt.Sprint("toJSON failed!:", err))
	}
	return bs, nil
}
