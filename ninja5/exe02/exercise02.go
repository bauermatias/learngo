package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	flavors   []string
}

func main() {

	p1 := person{
		firstName: "John",
		lastName:  "Doe",
		flavors:   []string{"Vanilla", "Chocolate", "Strawberry"},
	}

	p2 := person{
		firstName: "Jane",
		lastName:  "Doek",
		flavors:   []string{"Hazlenut", "Caramel", "Lemon"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.flavors {
			fmt.Println(i, val)
		}
		fmt.Println("-------")
	}
}
