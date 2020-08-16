package main

import "fmt"

/*
1.- Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
	42
	“James Bond”
	true
2.- Now print the values stored in those variables using
	a) a single print statement
	b) multiple print statements
*/

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("Single print:")
	fmt.Println(x, y, z)

	fmt.Println("Multiple print:")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
