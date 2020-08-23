package main

import "fmt"

func main() {
	p := func() {
		fmt.Println("This is my anonymous function")
	}

	p()
}
