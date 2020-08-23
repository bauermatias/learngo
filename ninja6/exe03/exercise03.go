package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "John",
		last:  "Doe",
		age:   42,
	}
	p1.speak()
}

func (p person) speak() {
	fmt.Println("Hi my name is", p.first, p.last, "and I am", p.age, "years old.")
}
