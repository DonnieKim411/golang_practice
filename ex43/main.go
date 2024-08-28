package main

import "fmt"

type engine struct {
	electric bool
}

type vehical struct {
	engine             engine
	make, model, color string
	doors              int
}

func main() {

	v1 := vehical{
		make:   "tesla",
		model:  "Y",
		color:  "white",
		doors:  4,
		engine: engine{electric: true},
	}

	v2 := vehical{
		make:   "toyota",
		model:  "corolla",
		color:  "gray",
		doors:  4,
		engine: engine{electric: false},
	}

	fmt.Println(v1, v2)

	fmt.Println(v1.engine)
	fmt.Println(v2.engine)
	fmt.Println(v1.engine.electric)
	fmt.Println(v2.engine.electric)

}
