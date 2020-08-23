package main

import "fmt"

func main() {
	rf := foo()
	rf()
}

func foo() func() {
	return func() {
		fmt.Println("This is a function that returns a function")
	}
}
