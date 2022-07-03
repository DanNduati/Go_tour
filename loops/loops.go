package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	// implement a square root function(Newton's method)

	z := float64(x / 2)

	var prev float64

	for {
		z -= (z*z - x) / (2 * z)
		if z == prev {
			break
		}
		//fmt.Println(z)
		prev = z
	}
	return z
}

func main() {
	x := float64(25)
	fmt.Printf("The squareroot of %v is %g", x, Sqrt(x))
}
