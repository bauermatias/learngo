package main

import "fmt"

/*
Write a program that
- assigns an int to a variable
- prints that int in decimal, binary, and hex
- shifts the bits of that int over 1 position to the left, and assigns that to a variable
- prints that variable in decimal, binary, and hex
*/

func main() {
	v := 12
	fmt.Printf("%d\n", v)
	fmt.Printf("%b\n", v)
	fmt.Printf("%x\n", v)
	y := v << 1
	fmt.Printf("%d\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x", y)
}
