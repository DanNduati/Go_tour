package main

import "fmt"

// The adder function return a closure.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// Go functions may be closures.
	// A closure is a function value that refrences values from outside its body
	// The function may access and assign to the refrenced variables; in this sense the function is "bound" to the variables
	// Each closure is bound to its own sum variable
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
