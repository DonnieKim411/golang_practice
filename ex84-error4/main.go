package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func test() sqrtError {
	return sqrtError{"foo", "bar", errors.New("this is an error")}
}

func main() {

	// a := test()
	_, err := sqrt(-10.23)
	fmt.Printf("%T\n", err)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		return 0, sqrtError{"hello", "world", fmt.Errorf("can't pass a value below 0: %v", f)}
	}
	return 42, nil
}
