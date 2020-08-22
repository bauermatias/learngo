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
		lastName:  "Doe",
		flavors:   []string{"Hazlenut", "Caramel", "Lemon"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for i, v := range p1.flavors {
		fmt.Println(i, v)
	}
}
