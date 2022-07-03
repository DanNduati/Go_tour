package main

import "fmt"

func main() {
	// Arrays
	// The type [n]T is an array of n values of type T
	// The expression var a [10]int
	// declares a variable a as an array of 10 integers
	var names [2]string
	names[0] = "Daniel"
	names[1] = "Chege"
	fmt.Println(names[0], names[1])
	fmt.Println(names)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
