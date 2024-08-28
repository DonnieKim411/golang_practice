package main

import "fmt"

func main() {
	sth := func() {
		fmt.Println("anonymous func!")
	}

	sth()
}
