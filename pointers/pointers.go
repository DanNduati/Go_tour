package main

import "fmt"

func main() {
	fmt.Println("Hello Pointers!")
	// Pointers hold the mem adderss of a variable/value
	// Declaration
	var p *int
	i := 42
	p = &i
	fmt.Println("The address for variable i is:", p)
	fmt.Println("The value in the adress is:", *p)
	*p = 21 //set value of i through pointer p
	fmt.Println("The new value of i is:", i)
}
