package main

import "fmt"

type customErr struct {
	info string
}

func (c customErr) Error() string {
	return fmt.Sprintf("Error Occured! %v", c.info)
}

func foo(e error) {
	fmt.Println("foo ran -", e)
}

func main() {

	ce := customErr{
		info: "hello error world!",
	}
	foo(ce)

}

// Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
// has a value of type error as a parameter. Create a value of type “customErr” and pass it into
// “foo”.
