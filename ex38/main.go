package main

import (
	"fmt"
)

func main() {
	m := make(map[string][]string)
	m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}
	m[`foo_bar`] = []string{`steaks`, `pizza`, `smoothie`}

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

}

/*
"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."
*/
