package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

type BySayings user

func (a BySayings) Len() int           { return len(a.Sayings) }
func (a BySayings) Swap(i, j int)      { a.Sayings[i], a.Sayings[j] = a.Sayings[j], a.Sayings[i] }
func (a BySayings) Less(i, j int) bool { return a.Sayings[i] < a.Sayings[j] }

func printAttribute(users []user) {
	for _, person := range users {
		fmt.Println(person.First, person.Last, person.Age)

		for _, saying := range person.Sayings {
			println(saying)
		}
		fmt.Println("-------")
	}
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// fmt.Println(users)
	printAttribute(users)

	// your code goes here
	// sort the []user by
	// ● age
	// ● last
	// Also sort each []string “Sayings” for each user
	// ● print everything in a way that is pleasant
	sort.Sort(ByAge(users))
	fmt.Println("----------------SORT BY AGE-------------")
	printAttribute(users)

	sort.Sort(ByLast(users))
	fmt.Println("----------------SORT BY Last-------------")
	printAttribute(users)

	for _, u := range users {
		sort.Sort(BySayings(u))
	}
	fmt.Println("----------------SORT EACH USER Saying-------------")
	printAttribute(users)

}
