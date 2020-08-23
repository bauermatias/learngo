package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6}
	n := foo(ii...)
	fmt.Println(n)

	jj := bar([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(jj)
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
