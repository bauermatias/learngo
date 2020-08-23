package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 15
}

func bar() (int, string) {
	return 12, "Twelve"
}
