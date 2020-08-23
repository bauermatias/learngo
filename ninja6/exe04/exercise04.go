package main

import "fmt"

func greet() {
	fmt.Println("it has been deferred and it's being called last")
}

func main() {

	defer greet()

	fmt.Println("The greet function was written before me, yet")

}
